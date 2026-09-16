# Софт для музыки на ноутбуке: плееры, кью-плееры, диджейские программы

Проверено по сайтам разработчиков 11.09.2026. Цены — без НДС, в валюте магазина разработчика.
Пометка «(по вторичным источникам)» — цена взята не с официального сайта, а из обзоров/магазинов.

Источник музыки: локальные файлы. Задачи: удобная навигация, кроссфейд между треками,
метки/точки внутри трека, кью-лист для сценария мероприятия.

---

## 1. Плееры для локальных файлов

Отсортировано: сначала бесплатные, затем платные.

### macOS + Windows

| Программа | Цена | Кроссфейд | Метки внутри трека | Заметки |
|---|---|---|---|---|
| **foobar2000** | Бесплатно (Win 2.x; mac 2.25.10) | Через DSP «Crossfader» (компонент) | Нет (только закладки) | Минималистичный, очень быстрый поиск и фильтры по библиотеке. Гибкая раскладка, плагины. Mac-версия проще Windows-версии |
| **VLC** | Бесплатно | Нет полноценного (только fade на паузе) | Закладки (Bookmarks) | Играет всё, но как музыкальный плеер с библиотекой слабый |
| **Strawberry** | Бесплатно на Linux; **macOS/Windows-сборки — только спонсорам, $7/мес** | Да, настраиваемый + gapless | Нет | Форк Clementine. Хорошая библиотека, умные плейлисты |
| **Dopamine 3** | Бесплатно (GPLv3, donationware), v3.0.10 | Да | Нет | Простой и красивый, Electron. Win/mac/Linux |

### Только macOS

| Программа | Цена | Кроссфейд | Метки | Заметки |
|---|---|---|---|---|
| **Музыка (Apple Music.app)** | Бесплатно, встроен | Да (Настройки → Воспроизведение → «Плавный переход») | Нет | Локальные файлы поддерживает. Для быстрого старта — достаточно |
| **Cog** | Бесплатно (GPL, форк Christopher Snowhill) | Нет | Нет | Легковесный, gapless, много редких форматов |
| **Colibri** | $19.99 (Mac App Store) | Нет (bit-perfect плеер) | Нет | Для lossless-коллекций, минимум интерфейса |
| **Doppler** | $30 разово, 7-дневный триал | Не подтверждено на сайте | Нет | Аккуратная библиотека, синхронизация с iPhone. macOS 11+ |
| **Swinsian** | $34.95 разово, есть триал | Да + gapless | Нет | Лучшая на mac навигация по большой библиотеке: колонки-браузеры, умные плейлисты, быстрый поиск. Замена iTunes «как раньше» |
| **VOX** | Бесплатно; Premium $4.99/мес или $49.99/год | Да | Нет | Premium в основном про облако. Реклама подписки навязчивая |
| **Audirvāna Origin** | $119.90 разово (по вторичным источникам; на сайте audirvana.com/price) | Нет | Нет | Аудиофильский плеер, не про навигацию по событию |

### Только Windows

| Программа | Цена | Кроссфейд | Метки | Заметки |
|---|---|---|---|---|
| **AIMP** | Бесплатно, v5.40 (сент. 2026); AIMP6 beta Win+Linux | Да + gapless | Закладки | Российская разработка, компактный, отличный эквалайзер и плейлисты. Русский интерфейс |
| **MusicBee** | Бесплатно, v3.6 (июнь 2026) | Да + gapless, отключение кроссфейда внутри альбома | Нет | Самая богатая библиотека/теги на Windows. Плагины, скины, горячие клавиши |
| **Winamp** | Бесплатно (5.9.2) | Через плагины | Нет | Классика, но развитие фактически остановлено |
| **MediaMonkey** | Бесплатно; Gold $39.95 (версия 2024.x) / $99.95 (пожизненно) | Да + авто-громкость | Нет | Для очень больших коллекций и синхронизации с устройствами |

**Итог по плеерам.** Метки внутри трека (cue points) обычные плееры почти не умеют — это уже задача
кью-плеера (раздел 2) или диджейского софта (раздел 3). Если нужен именно плеер:
на mac — Swinsian, на Windows — MusicBee или AIMP; foobar2000 — если нужна одна программа на обеих ОС.

---

## 2. Кью-плееры и саундборды (сценарий мероприятия)

Что это: список кью (треков/джинглов) в заданном порядке; нажатие «GO» запускает следующий,
с настраиваемыми fade in/out, точками старта/стопа внутри файла, авто-переходами и горячими клавишами.
Ровно то, что нужно для конференции или шоу. Отсортировано от простого к сложному.

| Программа | ОС | Цена | Уровень | Заметки |
|---|---|---|---|---|
| **Jingle Palette** / Jingle Palette Reloaded | Windows | Бесплатно | Простой | Сетка кнопок с джинглами, loop, нормализация. Для «нажал — сыграло». Из радийного мира |
| **Soundplant** | Win + mac, v59.5 (июнь 2026) | Бесплатно для некоммерческого; регистрация $60 (mp3/flac, выбор устройства вывода, MIDI, глобальные хоткеи) | Простой | Клавиатура = 72 кнопки со звуками. Быстро и надёжно, но без «списка кью» |
| **MultiPlay** | Windows | Бесплатно (в т.ч. коммерчески) | Простой–средний | Классический кью-плеер для театра/корпоративов. Последняя версия 2.5.5.0 — давно не обновляется |
| **Linux Show Player (LiSP)** | Linux (упомянут для полноты) | Бесплатно | Средний | На mac/Windows официально нет |
| **Farrago** (Rogue Amoeba) | macOS 14+ | $55 разово, триал = полная версия с ограничением | Средний | Саундборд с наборами, fade, hot keys, dual-output. Очень удобный UI, «для подкастов и живых мероприятий» |
| **Sports Sounds Pro** | Win + mac | $299.99/год (2 машины), триал 14 дней | Средний | Саундборд для спорта/ивентов: Stream Deck, таймер, запланированные кью |
| **SCS — Show Cue System** | Windows | Lite $62 · Standard $107 · Professional $161 · Pro Plus $214 · Platinum $268; временные лицензии от $21/30 дней; демо бесплатно | Средний–сложный | «QLab для Windows». Аудио + свет + MIDI, кью-лист с авто-переходами, точки старта/стопа в файле |
| **QLab 5** (Figure 53) | macOS | Бесплатно базовые аудио/видео/свет-кью без ограничения списка; платные лицензии — разово или аренда по дням (по вторичным источникам: от ~$4/день, постоянные $499–1199) | Сложный | Стандарт индустрии в театре/ивентах. Точные цены — в магазине qlab.app/shop (на странице цены не отдаются без JS) |
| **Ovation** (Merging) | Windows | Дорого, цену не публикуют | Очень сложный | Профессиональный медиа-сервер, избыточен |

**Итог по кью-плеерам.** Для конференции на mac: QLab (бесплатного режима хватит для запуска треков по кнопке)
или Farrago, если нужен «саундборд» с кнопками. На Windows: SCS Standard/Professional; бесплатно — MultiPlay.

---

## 3. Диджейский софт

Все программы — mac + Windows, если не указано иное. Отсортировано от простого к сложному
и от бесплатного к платному внутри уровня.

### Простой уровень (начать без контроллера)

| Программа | Цена | Заметки |
|---|---|---|
| **Mixxx 2.5.6** | Бесплатно, open source (Win 10 1809+, macOS 11+, Linux) | Полноценный: 4 деки, hot cues, лупы, авто-DJ с кроссфейдом, поддержка контроллеров. Лучший старт |
| **Serato DJ Lite** | Бесплатно | 2 деки, 4 hot cue, лупы 1–8 тактов, нет записи. Сэмплер только с поддерживаемым контроллером |
| **Cross DJ Free** (Mixvibes) | Бесплатно | Без ограничений по времени. Cross DJ 4 — $69.90, Pro 4 — $149.90 (по вторичным источникам) |
| **djay** (Algoriddim) | Бесплатно; PRO $6.99/мес или $49.99/год | Самый «дружелюбный» UI. Automix — автоматические переходы между треками, что близко к задаче «плеер с переходами» |
| **VirtualDJ** | Бесплатно дома без проф. оборудования; Home $4/мес; Pro $19/мес; Business $99/мес | Огромная база функций и контроллеров. Пожизненной Pro Infinity на странице цен больше нет — только подписки |
| **Traktor Play** (NI) | $49 разово | Упрощённый Traktor: 2 деки, stems, hot cues |

### Средний уровень (подготовка библиотеки, контроллеры)

| Программа | Цена | Заметки |
|---|---|---|
| **rekordbox** (AlphaTheta/Pioneer) | Free €0; Core €10/€12 мес; Creative €15/€18 мес; Professional €30/€36 мес (первая цена — при годовой оплате) | Бесплатного плана хватает для организации библиотеки и меток; платное — для контроллеров и эффектов |
| **Engine DJ Desktop** (Denon/inMusic) | Бесплатно, v5.0 (май 2026) | Подготовка библиотеки под железо Denon; как самостоятельный DJ-плеер ограничен |
| **Serato DJ Pro** | $11.99/мес или $299 разово; Suite $14.99/мес или $499 разово | Индустриальный стандарт для контроллеров; 8 hot cue, запись, stems |

### Сложный уровень

| Программа | Цена | Заметки |
|---|---|---|
| **Traktor Pro 4** (NI) | $149 разово | Stems, Flux, Remix Decks, глубокая настройка маппингов |
| **Ableton Live 12** | Intro ~$99; Standard ~$439; Suite ~$749 (по вторичным источникам; магазин Ableton для региона не отдаёт цены) | Не DJ-софт, а DAW с сессионным видом: клипы, сцены, переходы по кнопке, автоматизация. Максимальная гибкость и максимальный порог входа. Live Lite бесплатно в комплекте с некоторым железом |

**Итог по DJ.** Начать — Mixxx (бесплатно, всё есть) или djay (самый простой). Если понадобятся контроллеры
и «взрослая» библиотека — rekordbox Free → Serato DJ Pro / Traktor Pro 4.

---

## Что не удалось подтвердить на официальных сайтах

- QLab 5: точные цены лицензий (магазин отдаёт их только через JS).
- Ableton Live 12: магазин не показывает цены для региона.
- Audirvāna Origin, Cross DJ 4, Colibri: цены взяты из обзоров/магазинов, не с главной страницы разработчика.
- Doppler: поддержка кроссфейда на сайте не описана.

---

## 4. Открытый софт на GitHub

Поиск по GitHub Search API 11.09.2026 (запросы: music player, soundboard, cue player, show control,
dj, radio automation, cartwall и др.). Звёзды и дата последнего коммита — на момент поиска.
Платформы и функции сверены по README репозиториев. Отсортировано по звёздам внутри группы.

Исключены: клиенты стриминга (YesPlayMusic, nuclear, Feishin, SPlayer — нужен сервер или онлайн-сервис),
Android-плееры, библиотеки для разработчиков.

### 4.1 Плееры для локальной библиотеки

| Репозиторий | ★ | Коммит | ОС | Стек | Лицензия | Кроссфейд | Заметки |
|---|---|---|---|---|---|---|---|
| [mixxxdj/mixxx](https://github.com/mixxxdj/mixxx) | 7135 | 2026-09 | Win/mac/Linux | C++/Qt | GPL-2 | Да (Auto DJ) | См. раздел 3 — единственный полноценный open-source DJ |
| [harmonoid/harmonoid](https://github.com/harmonoid/harmonoid) | 4731 | 2026-09 | Win/mac/Linux | Dart/Flutter, mpv | **PolyForm Strict — не свободная** | Нет данных | Gapless, теги, тексты. Красивый Material 3 |
| [clementine-player/Clementine](https://github.com/clementine-player/Clementine) | 4249 | 2026-09 | Win/mac/Linux | C++/Qt | GPL-3 | Да | Классика; развитие вялое, официальные релизы редкие. Актуальный форк — Strawberry |
| [strawberrymusicplayer/strawberry](https://github.com/strawberrymusicplayer/strawberry) | 3953 | 2026-09 | Win/mac/Linux | C++/Qt | GPL-3 | Да + gapless | Официальные бинарники Win/mac только спонсорам ($7/мес), но исходники открыты: есть сторонние сборки [TheFreeman193/StrawberryWindows](https://github.com/TheFreeman193/StrawberryWindows) и [stm7128/strawberry-autobuild-windows](https://github.com/stm7128/strawberry-autobuild-windows) |
| [Taiko2k/Tauon](https://github.com/Taiko2k/Tauon) | 2848 | 2026-09 | Win/mac/Linux | Python/SDL | GPL-3 | Нет данных в README | Gapless, CUE-листы, быстрый UI, DMG для mac есть в релизах |
| [digimezzo/dopamine](https://github.com/digimezzo/dopamine) | 2345 | 2026-09 | Win/mac/Linux | TS/Electron | GPL-3 | Да | См. раздел 1 |
| [martpie/museeks](https://github.com/martpie/museeks) | 2136 | 2026-07 | Win/mac/Linux | Rust/Tauri + React | MIT | Нет | Намеренно простой. Плейлисты, m3u, скорость воспроизведения |
| [quodlibet/quodlibet](https://github.com/quodlibet/quodlibet) | 1751 | 2026-08 | Win/mac/Linux | Python/GTK | GPL-2 | Через плагин | 90+ плагинов, сильный поиск по тегам, теггер Ex Falso |
| [audacious-media-player/audacious](https://github.com/audacious-media-player/audacious) | 1250 | 2026-08 | Win/mac/Linux/BSD | C++/Qt или GTK | BSD-2 | Плагин Crossfade | Лёгкий, скины Winamp 2, gapless, ReplayGain. На mac ставится через Homebrew |
| [Sandakan/Nora](https://github.com/Sandakan/Nora) | 1215 | 2026-09 | Win/mac/Linux | TS/Electron | MIT | Нет данных | Плейлисты, синхронные тексты, редактор тегов |
| [basharovV/musicat](https://github.com/basharovV/musicat) | 937 | 2026-08 | Win/mac/Linux | Svelte/Tauri | GPL-3 | Нет (gapless есть) | Умные плейлисты, теггер, waveform, статистика |
| [Anthonyy232/Nagi](https://github.com/Anthonyy232/Nagi) | 294 | 2026-09 | **Windows** | C#/WinUI 3, libVLC | GPL-3 | Нет данных | Современный Fluent UI, умные плейлисты, ReplayGain, эквалайзер |
| [stsaz/phiola](https://github.com/stsaz/phiola) | 251 | 2026-09 | Win/mac/Linux | C | BSD-2 | Нет данных | Мгновенный старт, портативный; GUI только на Win/Linux, на mac — CLI/TUI |
| [zhongyang219/MusicPlayer2](https://github.com/zhongyang219/MusicPlayer2) | 6659 | 2026-09 | **Windows** | C++/MFC, BASS | GPL-3 | Нет данных | Китайский проект с английским UI; тексты, эквалайзер, CUE |
| [jfchapman/VUPlayer](https://github.com/jfchapman/VUPlayer) | 105 | 2026-07 | **Windows** | C | MIT | Нет данных (gapless есть) | Библиотека в SQLite, MusicBrainz, Last.fm |
| [rvarunrathod/HiFidelity](https://github.com/rvarunrathod/HiFidelity) | 312 | 2026-01 | **macOS 14+** | Swift, BASS | не указана | Нет | Bit-perfect, hog mode, gapless, EBU R128 |
| [losnoco/Cog](https://github.com/losnoco/Cog) | — | 2026 | **macOS** | Obj-C | GPL-2 | Нет | См. раздел 1 (форк Snowhill) |
| [kartik-venugopal/aural-player](https://github.com/kartik-venugopal/aural-player) | 1074 | архив 2025-06 | **macOS** | Swift | MIT | — | Winamp-подобный; **заархивирован**, но бинарник скачать можно |
| [reupen/columns_ui](https://github.com/reupen/columns_ui) | 622 | 2026-09 | Windows | C++ | — | — | Не плеер, а альтернативный UI для foobar2000 |

Не открытый код, хотя лежит на GitHub: [Moriafly/SaltPlayerSource](https://github.com/Moriafly/SaltPlayerSource) (7342 ★) —
только трекер задач и релизы.

### 4.2 Кью-плееры и саундборды

| Репозиторий | ★ | Коммит | ОС | Стек | Лицензия | Что умеет |
|---|---|---|---|---|---|---|
| [Soundux/Soundux](https://github.com/Soundux/Soundux) | 2042 | 2026-01 | Win/Linux (**mac нет**) | C++ | GPL-3 | Саундборд: сетка/список, хоткеи, вывод в конкретное приложение (VB-Cable на Win). Идёт полный rewrite |
| [FrancescoCeruti/linux-show-player](https://github.com/FrancescoCeruti/linux-show-player) | 308 | 2026-08 | **только Linux** | Python/PyQt5, GStreamer | GPL-3 | Кью-плеер для театра, MIDI/OSC, таймкод. Для mac/Win официально не собирается |
| [dan0v/AmplitudeSoundboard](https://github.com/dan0v/AmplitudeSoundboard) | 164 | 2026-05 | Win/mac/Linux | C#/Avalonia | GPL-3 | Сетка клипов, хоткеи, громкость и устройство вывода на каждый клип |
| [space928/QPlayer](https://github.com/space928/QPlayer) | 47 | 2026-09 | **Windows** | C#/WPF, NAudio | GPL-3 | Кью-лист для театра: fade in/out, пре-делей, EQ на кью, лимитер, OSC |
| [tdoukinitsas/liveplay](https://github.com/tdoukinitsas/liveplay) | 42 | 2026-09 | Win/mac/Linux | C++20 + Electron/Vue | AGPL-3 | Ближе всего к QLab: кью-лист с группами, fade, авто-переход, сетка «GO», MIDI, хоткеи, LUFS-метры, маршрутизация на несколько карт, LTC, удалённое управление по LAN |
| [suterma/replayer-pwa](https://github.com/suterma/replayer-pwa) | 20 | 2026-09 | Браузер (PWA), любая ОС | JS/Vue | AGPL-3 | Именно «метки внутри трека»: именованные точки с хоткеями, наборы треков, Bluetooth-нампад как пульт. Fade нет. Работает на web.replayer.app |
| [stoatworks-labs/simplecue](https://github.com/stoatworks-labs/simplecue) | 3 | 2026-09 | Win/mac/Linux | C++20/JUCE | MIT (бинарники AGPL) | Кью с in/out-точками, 5 кривых fade, связи между кью (мгновенно/по концу/кроссфейд), лупы, vamp, OSC/MIDI/MSC. v0.4.1, «на живом шоу ещё не запускался» |
| [wb56/deckrelay](https://github.com/wb56/deckrelay) | 1 | 2026-09 | **Windows** | Python + VLC/FFmpeg | open source | «Плеер для мероприятия»: две деки, автоматические кроссфейды, cue points, очередь, защита от клиппинга. v1.0, v2 beta |
| [SanderKlootwijk/cartwall](https://github.com/SanderKlootwijk/cartwall) | 0 | 2026-05 | Linux (QML) | QML | — | Сетка кнопок с джинглами, свежий проект |
| [aerialsbroadcast/RapidFire](https://github.com/aerialsbroadcast/RapidFire) | 0 | 2026-08 | **Windows** | — | **закрытый код**, бесплатно | Картволл на 1536 пэдов, задержка ~3 мс, fade, лог эфира. Только бинарник |

Мелкие/неактивные: [eeeeeta/sqa](https://github.com/eeeeeta/sqa) (Rust, «QLab alternative», заброшен 2017),
[evandelisle/showq](https://github.com/evandelisle/showq) (C++, 2023), [micahmo/SoundBoard](https://github.com/micahmo/SoundBoard) (Win, 2023),
[viluon/afx](https://github.com/viluon/afx) (Rust, кроссплатформенный саундборд, ранняя стадия),
[cg2121/obs-soundboard](https://github.com/cg2121/obs-soundboard) (док саундборда внутри OBS — если звук идёт через OBS).

### 4.3 DJ и авто-микс

| Репозиторий | ★ | Коммит | ОС | Стек | Лицензия | Заметки |
|---|---|---|---|---|---|---|
| [mixxxdj/mixxx](https://github.com/mixxxdj/mixxx) | 7135 | 2026-09 | Win/mac/Linux | C++/Qt | GPL-2 | Единственный зрелый open-source DJ-софт. Auto DJ с кроссфейдом — готовый «плеер с переходами» |
| [dylanljones/pyrekordbox](https://github.com/dylanljones/pyrekordbox) | 478 | 2026-09 | — | Python | MIT | Не плеер: работа с базой rekordbox (экспорт меток, плейлистов) |
| [digital-dj-tools/dj-data-converter](https://github.com/digital-dj-tools/dj-data-converter) | 218 | 2025-06 | CLI | Clojure | — | Конвертер библиотек между Traktor/rekordbox/Serato |
| [walkywalker/automix](https://github.com/walkywalker/automix) | 24 | 2023 | — | C++ | — | Автоматическое DJ-сведение, эксперимент |
| [fabriziosalmi/mixi](https://github.com/fabriziosalmi/mixi) | 6 | 2026-09 | Браузер | TS + Rust/Wasm | — | Две деки в браузере, WebGPU-визуал. Ранняя стадия |
| [waypostmaster/deckwave](https://github.com/waypostmaster/deckwave) | 5 | 2026-09 | Браузер | JS | — | Анализ библиотеки и сборка beat-locked сета в браузере |
| [sebedard-creator/MixCanvas](https://github.com/sebedard-creator/MixCanvas) | 1 | 2026-09 | — | Rust | — | Сведение миксов на таймлайне, офлайн. Совсем новый |

Радиоавтоматика (эфирный плейаут, избыточно для конференции, но умеет джинглы и переходы по расписанию):
[ElvishArtisan/rivendell](https://github.com/ElvishArtisan/rivendell) (242 ★, только Linux),
[libretime/libretime](https://github.com/libretime/libretime) (941 ★, серверный, web),
[openbroadcaster/obplayer](https://github.com/openbroadcaster/obplayer) (143 ★, Linux).

### Итог по GitHub

- **Плеер под mac и Windows:** Strawberry (собрать самому или взять стороннюю сборку), Tauon, Audacious, Quod Libet.
  Из «молодых» самый живой — musicat (Tauri, умные плейлисты).
- **Кью-плеер под задачу мероприятия:** liveplay — единственный open-source, который закрывает кью-лист + fade +
  авто-переходы + сетку GO на обеих ОС. QPlayer — если только Windows. simplecue — интересен, но сырой.
- **Метки внутри трека без установки:** replayer-pwa прямо в браузере.
- **DJ:** только Mixxx; остальное — эксперименты и утилиты.
