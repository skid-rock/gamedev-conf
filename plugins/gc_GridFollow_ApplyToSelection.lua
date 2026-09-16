-- gc_GridFollow_ApplyToSelection.lua
-- Разово перестроить сетку под выделенный айтем или дорожку.
-- Нужен, когда фоновый режим выключен, а сетку подогнать надо.

local info = debug.getinfo(1, "S").source:sub(2)
local dir = info:match("^(.*)[/\\]")
local lib = dofile(dir .. "/gc_gridfollow_lib.lua")

local item = lib.resolveTargetItem()
local ok, message = lib.applyGrid(item)

if not ok then
  reaper.MB(message or "не удалось построить сетку", "GridFollow", 0)
end
