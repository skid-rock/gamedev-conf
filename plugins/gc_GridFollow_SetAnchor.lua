-- gc_GridFollow_SetAnchor.lua
-- Задать темп и первую долю для айтема под курсором.
-- Делается один раз на трек: дальше сетка строится по этим данным.
--
-- Порядок: поставь курсор редактирования ровно на первую долю трека,
-- выдели айтем (или его дорожку) и запусти это действие.

local info = debug.getinfo(1, "S").source:sub(2)
local dir = info:match("^(.*)[/\\]")
local lib = dofile(dir .. "/gc_gridfollow_lib.lua")

local item = lib.resolveTargetItem()
if not item then
  reaper.MB("Не найден айтем. Выдели айтем или дорожку с аудио.", "GridFollow", 0)
  return
end

local position = reaper.GetMediaItemInfo_Value(item, "D_POSITION")
local length = reaper.GetMediaItemInfo_Value(item, "D_LENGTH")
local cursor = reaper.GetCursorPosition()

-- Якорь берём из курсора, если тот стоит внутри айтема. Иначе — ранее сохранённый.
local anchorOffset
if cursor >= position and cursor <= position + length then
  anchorOffset = cursor - position
else
  anchorOffset = lib.getAnchorOffset(item)
end

local currentBpm, source = lib.getItemBpm(item)
local defaults = string.format("%s,%.4f",
  currentBpm and string.format("%.4f", currentBpm) or "",
  anchorOffset)

local ok, input = reaper.GetUserInputs(
  "GridFollow: темп и первая доля",
  2,
  "BPM (слышимый),Якорь: секунд от начала айтема,extrawidth=80",
  defaults)

if not ok then return end

local bpmText, anchorText = input:match("([^,]*),([^,]*)")
local bpm = tonumber(bpmText)
local anchor = tonumber(anchorText)

if not bpm or bpm <= 0 then
  reaper.MB("BPM задан неверно.", "GridFollow", 0)
  return
end
if not anchor then anchor = 0 end

reaper.Undo_BeginBlock()
lib.setItemGridData(item, bpm, anchor)
local applied, message = lib.applyGrid(item)
reaper.Undo_EndBlock("GridFollow: задать темп и якорь", -1)

if applied then
  reaper.ShowConsoleMsg(string.format(
    "GridFollow: %s | источник темпа до правки: %s\n", message, source or "не найден"))
else
  reaper.MB(message or "не удалось построить сетку", "GridFollow", 0)
end
