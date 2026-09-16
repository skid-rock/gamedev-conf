-- gc_gridfollow_lib.lua
-- Общая библиотека набора «сетка следует за дорожкой» для REAPER.
-- Подключается из скриптов-действий через dofile().
--
-- Идея: темпо-карта в REAPER одна на проект и живёт во времени, поэтому
-- «своя сетка у каждой дорожки» эмулируется так — при переключении цели
-- вся темпо-карта переписывается под выбранный айтем.
-- Обязательное условие: таймбейс всех дорожек = Time, иначе смена темпа
-- потащит за собой аудио и огибающие.

local M = {}

-- Ключи расширенных данных айтема. Живут внутри проекта и переживают сохранение.
M.EXT_BPM    = "P_EXT:gc_grid_bpm"     -- слышимый темп айтема (уже с учётом playrate)
M.EXT_ANCHOR = "P_EXT:gc_grid_anchor"  -- смещение первой доли от начала айтема, в секундах

-- Секция ExtState для связи между экземплярами фонового скрипта.
M.EXT_SECTION = "gc_gridfollow"

-- Разумные границы, в которые обязан попадать темп при автоопределении.
M.BPM_MIN = 40
M.BPM_MAX = 300

--- Печать в консоль ReaScript.
function M.msg(text)
  reaper.ShowConsoleMsg(tostring(text) .. "\n")
end

--- Вытащить первое похожее на темп число из строки.
-- Отсеивает номера дорожек и годы: берутся только значения в BPM_MIN..BPM_MAX.
function M.parseBpm(text)
  if not text or text == "" then return nil end
  for candidate in text:gmatch("%d+%.?%d*") do
    local value = tonumber(candidate)
    if value and value >= M.BPM_MIN and value <= M.BPM_MAX then
      return value
    end
  end
  return nil
end

--- Прочитать темп из метаданных исходного файла.
-- Конкретные идентификаторы зависят от формата (acid-чанк в WAV, теги в MP3),
-- поэтому сначала запрашиваем весь список идентификаторов и ищем в нём
-- подходящий по имени — так не приходится угадывать ключ.
function M.bpmFromMetadata(take)
  if not take or reaper.TakeIsMIDI(take) then return nil end
  local source = reaper.GetMediaItemTake_Source(take)
  if not source then return nil end

  local ok, list = reaper.GetMediaFileMetadata(source, "")
  if not ok or ok == 0 or not list or list == "" then return nil end

  for identifier in list:gmatch("[^\r\n]+") do
    local lowered = identifier:lower()
    if lowered:find("bpm") or lowered:find("tempo") then
      local got, value = reaper.GetMediaFileMetadata(source, identifier)
      if got and got ~= 0 then
        local bpm = M.parseBpm(value)
        if bpm then return bpm end
      end
    end
  end
  return nil
end

--- Имя активного тейка айтема.
local function takeName(item)
  local take = reaper.GetActiveTake(item)
  if not take then return nil, nil end
  local _, name = reaper.GetSetMediaItemTakeInfo_String(take, "P_NAME", "", false)
  return name, take
end

--- Определить темп айтема.
-- Порядок источников: сохранённое значение → имя тейка → имя дорожки → метаданные файла.
-- Возвращает bpm и метку источника, либо nil.
function M.getItemBpm(item)
  if not item then return nil, nil end

  local ok, stored = reaper.GetSetMediaItemInfo_String(item, M.EXT_BPM, "", false)
  if ok and stored ~= "" then
    local bpm = tonumber(stored)
    if bpm then return bpm, "сохранён" end
  end

  local name, take = takeName(item)
  local bpm = M.parseBpm(name)
  if bpm then return bpm, "имя тейка" end

  local track = reaper.GetMediaItem_Track(item)
  if track then
    local _, trackName = reaper.GetSetMediaTrackInfo_String(track, "P_NAME", "", false)
    bpm = M.parseBpm(trackName)
    if bpm then return bpm, "имя дорожки" end
  end

  bpm = M.bpmFromMetadata(take)
  if bpm then
    -- В метаданных лежит темп исходника; слышимый темп зависит от растяжения.
    local rate = take and reaper.GetMediaItemTakeInfo_Value(take, "D_PLAYRATE") or 1.0
    if rate and rate > 0 then bpm = bpm * rate end
    return bpm, "метаданные файла"
  end

  return nil, nil
end

--- Смещение первой доли от начала айтема, в секундах.
function M.getAnchorOffset(item)
  if not item then return 0 end
  local ok, stored = reaper.GetSetMediaItemInfo_String(item, M.EXT_ANCHOR, "", false)
  if ok and stored ~= "" then
    local value = tonumber(stored)
    if value then return value end
  end
  return 0
end

--- Записать темп и якорь в расширенные данные айтема.
function M.setItemGridData(item, bpm, anchorOffset)
  if not item then return false end
  reaper.GetSetMediaItemInfo_String(item, M.EXT_BPM, tostring(bpm), true)
  reaper.GetSetMediaItemInfo_String(item, M.EXT_ANCHOR, tostring(anchorOffset), true)
  return true
end

--- Найти айтем, под который строить сетку.
-- Приоритет: выделенный айтем → айтем выделенной дорожки под курсором → ближайший к курсору.
function M.resolveTargetItem()
  local selected = reaper.GetSelectedMediaItem(0, 0)
  if selected then return selected end

  local track = reaper.GetSelectedTrack(0, 0)
  if not track then return nil end

  local cursor = reaper.GetCursorPosition()
  local nearest, nearestDistance = nil, nil

  for i = 0, reaper.CountTrackMediaItems(track) - 1 do
    local item = reaper.GetTrackMediaItem(track, i)
    local position = reaper.GetMediaItemInfo_Value(item, "D_POSITION")
    local length = reaper.GetMediaItemInfo_Value(item, "D_LENGTH")

    if cursor >= position and cursor <= position + length then
      return item
    end

    local distance = math.min(math.abs(cursor - position), math.abs(cursor - (position + length)))
    if not nearestDistance or distance < nearestDistance then
      nearest, nearestDistance = item, distance
    end
  end

  return nearest
end

--- Перестроить темпо-карту проекта под один айтем.
-- Возвращает true при успехе, либо false и причину.
function M.applyGrid(item)
  if not item then return false, "не выбрана дорожка или айтем" end

  local bpm, source = M.getItemBpm(item)
  if not bpm then
    return false, "у айтема не задан BPM — сначала запусти SetAnchor"
  end

  local position = reaper.GetMediaItemInfo_Value(item, "D_POSITION")
  local anchor = position + M.getAnchorOffset(item)
  if anchor < 0 then anchor = 0 end

  reaper.PreventUIRefresh(1)

  -- Сносим всю текущую карту: сетка проекта временно принадлежит одному айтему.
  while reaper.CountTempoTimeSigMarkers(0) > 0 do
    reaper.DeleteTempoTimeSigMarker(0, 0)
  end

  -- Базовый темп проекта — тот же, чтобы слева от якоря шаг сетки не отличался.
  reaper.SetCurrentBPM(0, bpm, false)

  if anchor > 0 then
    reaper.AddTempoTimeSigMarker(0, anchor, bpm, 4, 4, false)
    local index = reaper.FindTempoTimeSigMarker(0, anchor)
    if index and index >= 0 then
      -- Флаги: 1 — начать новый такт, 4 — разрешить неполный такт слева,
      -- 16 — сбросить сетку линейки. Сумма вместо побитового ИЛИ ради
      -- совместимости со старым синтаксисом Lua.
      reaper.GetSetTempoTimeSigMarkerFlag(0, index, 1 + 4 + 16, true)

      -- Маркер мог прилипнуть к сетке прежнего темпа — возвращаем ровно на якорь.
      local ok, markerPos = reaper.GetTempoTimeSigMarker(0, index)
      if ok and math.abs(markerPos - anchor) > 1e-6 then
        reaper.SetTempoTimeSigMarker(0, index, anchor, -1, -1, bpm, 4, 4, false)
      end
    end
  end

  reaper.PreventUIRefresh(-1)
  reaper.UpdateTimeline()
  reaper.UpdateArrange()

  return true, string.format("%.3f BPM (%s), якорь %.3f с", bpm, source or "?", anchor)
end

--- Шаг сетки в тактах 4/4. В API деление задаётся в долях целой ноты: 0.25 — четверть.
function M.setGridBars(bars)
  if not bars or bars <= 0 then return end
  reaper.GetSetProjectGrid(0, true, bars * 1.0)
end

--- Перевести дорожки и айтемы на таймбейс Time.
-- Без этого смена темпа будет двигать и растягивать материал.
function M.forceTimeTimebase()
  local tracks, items = 0, 0

  for i = 0, reaper.CountTracks(0) - 1 do
    local track = reaper.GetTrack(0, i)
    -- 0 = time для дорожки
    reaper.SetMediaTrackInfo_Value(track, "C_BEATATTACHMODE", 0)
    tracks = tracks + 1
  end

  for i = 0, reaper.CountMediaItems(0) - 1 do
    local item = reaper.GetMediaItem(0, i)
    -- -1 = наследовать таймбейс дорожки; собственного значения «time» у айтема нет
    reaper.SetMediaItemInfo_Value(item, "C_BEATATTACHMODE", -1)
    reaper.SetMediaItemInfo_Value(item, "C_AUTOSTRETCH", 0)
    items = items + 1
  end

  return tracks, items
end

return M
