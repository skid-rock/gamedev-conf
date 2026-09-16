-- gc_GridFollow_Toggle.lua
-- Фоновый режим: сетка сама перестраивается под выделенную дорожку или айтем.
-- Повторный запуск действия выключает режим.
--
-- Перед первым использованием запусти gc_GridFollow_PrepareProject.lua,
-- иначе смена темпа потащит за собой материал.

local info = debug.getinfo(1, "S").source:sub(2)
local dir = info:match("^(.*)[/\\]")
local lib = dofile(dir .. "/gc_gridfollow_lib.lua")

local RUNNING_KEY = "running"
local POLL_INTERVAL = 0.15  -- как часто опрашивать выделение, в секундах

local _, _, sectionID, commandID = reaper.get_action_context()

-- Второй запуск гасит уже работающий экземпляр: тот увидит снятый флаг и выйдет.
if reaper.GetExtState(lib.EXT_SECTION, RUNNING_KEY) == "1" then
  reaper.SetExtState(lib.EXT_SECTION, RUNNING_KEY, "0", false)
  return
end

reaper.SetExtState(lib.EXT_SECTION, RUNNING_KEY, "1", false)
reaper.SetToggleCommandState(sectionID, commandID, 1)
reaper.RefreshToolbar2(sectionID, commandID)

reaper.atexit(function()
  reaper.SetExtState(lib.EXT_SECTION, RUNNING_KEY, "0", false)
  reaper.SetToggleCommandState(sectionID, commandID, 0)
  reaper.RefreshToolbar2(sectionID, commandID)
end)

local lastKey = nil
local nextPoll = 0

--- Слепок состояния цели: меняется при смене айтема, его сдвиге или правке данных.
local function stateKey(item)
  if not item then return "" end
  local position = reaper.GetMediaItemInfo_Value(item, "D_POSITION")
  local bpm = lib.getItemBpm(item) or 0
  local anchor = lib.getAnchorOffset(item)
  return string.format("%s|%.6f|%.6f|%.6f", tostring(item), position, bpm, anchor)
end

local function loop()
  if reaper.GetExtState(lib.EXT_SECTION, RUNNING_KEY) ~= "1" then
    return  -- выключили вторым запуском действия
  end

  local now = reaper.time_precise()
  if now >= nextPoll then
    nextPoll = now + POLL_INTERVAL

    local item = lib.resolveTargetItem()
    if item and not reaper.ValidatePtr2(0, item, "MediaItem*") then
      item = nil
    end

    local key = stateKey(item)
    if key ~= lastKey then
      lastKey = key
      -- Молча пропускаем айтемы без темпа: фоновый режим не должен сыпать диалогами.
      if item then lib.applyGrid(item) end
    end
  end

  reaper.defer(loop)
end

loop()
