-- gc_GridFollow_PrepareProject.lua
-- Подготовить проект: перевести дорожки и айтемы на таймбейс Time
-- и выставить шаг сетки в тактах.
--
-- Без таймбейса Time переключение сетки будет двигать и растягивать материал —
-- запусти это действие один раз на проект, до работы с GridFollow.

local info = debug.getinfo(1, "S").source:sub(2)
local dir = info:match("^(.*)[/\\]")
local lib = dofile(dir .. "/gc_gridfollow_lib.lua")

local ok, input = reaper.GetUserInputs(
  "GridFollow: подготовка проекта",
  1,
  "Шаг сетки в тактах (1, 2, 4...)",
  "1")
if not ok then return end

local bars = tonumber(input)
if not bars or bars <= 0 then bars = 1 end

reaper.Undo_BeginBlock()
local tracks, items = lib.forceTimeTimebase()
lib.setGridBars(bars)
reaper.Undo_EndBlock("GridFollow: подготовка проекта", -1)

reaper.UpdateArrange()
reaper.MB(string.format(
  "Таймбейс Time выставлен.\nДорожек: %d\nАйтемов: %d\nШаг сетки: %d такт(а/ов)",
  tracks, items, bars), "GridFollow", 0)
