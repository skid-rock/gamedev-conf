# Альбом на базе jump-up — план, промпты, журнал

Продолжение файла 04. Там — поджанры DnB по отдельности. Здесь — один альбом на 10–11 треков, где jump-up — база, а
остальное — соседи. Решения приняты 12 сентября 2026 на grill-сессии, дедлайн — безлимит Suno до 13 сентября, на всё
около 21 часа.

Правила промптов — из файла 02. Что уже проверено на слух — треки в `gamepush/`: там только понравившееся, все на v6
(chirp-hawk). Контрольная точка альбома — трек «Jump-up 8-bit» (id `7fb8ea93-43c5-4149-a2a9-68573b357a6f`), его Style
воспроизведён ниже как J1 слово в слово.

## Что решили

- **Альбом самостоятельный.** Мотив конференции не тащим, Cover-цепочка из файла 03 не нужна. Треки генерируем массово и
  отбираем на вкус.
- **Jump-up — база, соседи — разнообразие.** Шесть треков ядра, по одному-два трека jungle, dancefloor, halftime,
  liquid-лайт, плюс breakcore-вставка как бонус. Neurofunk — только приправа внутри jump-up, отдельным треком не идёт.
  Techstep, sambass, atmospheric, minimal — нет.
- **Игровая тема — через палитру звука.** Концепт «прохождение игры» отклонён. Никаких каверов игровых мелодий: права у
  Konami, Sega и прочих. Единственная мелодия без риска — «Коробейники» из Tetris, народная. Ощущение «знакомой музыки
  из игры» передаём идиомами эпохи (раздел «Знакомость без нарушения прав»).
- **Модель v6.** v6-wild пробовали, не понравилось. Не трогаем.
- **Хук в ядре — бас, лид его дублирует.** Формула из J1. В 3–4 соседях пробуем обратное: лид ведёт свою мелодию, бас
  держит грув.
- **Названия — в конце.** В Suno при генерации — рабочие коды (J1, N2, R3). В финальных названиях нет имён игр и
  персонажей: торговые марки.
- **Порядок — энергетическая дуга,** предлагаю я, правится на слух.
- Финальный формат, WAV и нормализация громкости — отдельная забота, к этому файлу не относится.

## Инварианты — что склеивает альбом

Есть в каждом треке без исключения:

1. **8-bit или FM-лид** — слышимый голос сверху. В ядре дублирует бас, у соседей может вести мелодию.
2. **Словарь аркадных SFX** — монетки, прыжки, взрывы, лазеры, «stage clear». Хотя бы один вид на трек, в интро или как
   филлы.
3. **Loud club mix** — одна и та же формулировка микса в каждом Style, чтобы треки сидели на одном уровне давления.

Коридор, из которого можно выйти один-два раза: темп 170–176 (исключения — breakcore-вставка и, возможно, halftime).

Ось разнообразия — **барабаны**: two-step в ядре, amen в jungle, ровный роллер в dancefloor, half-time в одном треке.

## Знакомость без нарушения прав

Два приёма, оба проверяются в залпе 1:

- **R1 — названия игр в Style.** «Contra NES soundtrack style» и подобное. По сообщениям сообщества Suno имена артистов
  режет, названия игр — [не проверено]. Если модель их читает и звук сдвигается — используем точечно. Имена композиторов
  не пишем: даже если пройдут, не поймём, что сработало.
- **R2 — идиомы эпохи вместо названий.** NES-арпеджио аккордов, гармонический минор, хроматический пробег перед дропом,
  фрагмент фанфары «stage clear», звуки «1-up» и монетки как филлы, FM slap bass, стаккато-марш. Безопасно и, скорее
  всего, надёжнее, чем названия.

Юридическая граница: палитра и приёмы — можно, мелодии — нет. Cover R3 делается из собственного трека, поэтому чист.

**Итог после залпа 1 (12.09.2026):** названия игр закрыты. R1 сдвинул звук в обеих парах, но не к Contra, а в случайную сторону (NES-нуар). Имя работает как толчок, не как прицел. Идиомы R2 дали «похоже на игру» надёжнее. Дальше — только консоль и настроение словами, без имён игр.

## Настройки

| Что             | Ядро (J)                                    | Соседи (N) | Разведка (R)                                   |
|-----------------|---------------------------------------------|------------|------------------------------------------------|
| Режим           | Custom, Instrumental вкл                    | то же      | то же                                          |
| Style Influence | 80%                                         | 80%        | 80%                                            |
| Weirdness       | 45%                                         | 35%        | 45% (R3 — 35%)                                 |
| Variety         | Normal                                      | Normal     | 0 — иначе разброс дублей забьёт эффект промпта |
| Max Mode        | выкл                                        | выкл       | R3 (Cover) — вкл                               |
| Duration Custom | 3:15                                        | 3:45       | как у базового промпта                         |
| Exclude Styles | `lyrics, sung words, rap, vocals, singing, soft pads, piano, brass, trumpets, horns, fanfare, orchestral, strings, live instruments, lead guitar, guitar solo, rock band` | то же плюс по промпту | по промпту |

Слово «instrumental» и «no ...» в Style не пишем — тумблер и Exclude. Формат Style — как в файле 04. Suno у себя
переписывает промпт в форму с двоеточием (так он выглядит в json из gamepush), это не наша
забота [допущение: в json сохранён переписанный текст, а вводилась форма файла 04].

## План на 21 час — три залпа

| Залп | Что                                                      | Пар   | Когда                |
|------|----------------------------------------------------------|-------|----------------------|
| 1    | R1, R2, R3 по одной паре + J1–J6 по две пары             | 15    | сразу                |
| 2    | J1–J6 по две пары (второй заход, раздел ниже) + N1–N5 по две пары | 22    | после оценок залпа 1 |
| 3    | добор дыр в трек-листе                                   | 4–6   | после оценок залпа 2 |

После залпа 3: отбор 10–11, порядок, названия.

Оценки — в чат, одной строкой на трек: `код — оценка 1–5 — что не так`. Понравившееся — качать в `gamepush/` через
suno-download с `--name '{artist} - {title}'`, где title — рабочий код. Журнал в конце файла веду я.

## Целевой трек-лист (план, не порядок)

| Код | Что                        | Поджанр     | Ось разнообразия                                                    |
|-----|----------------------------|-------------|---------------------------------------------------------------------|
| J1  | NES, наглый                | jump-up     | pulse-лид дублирует talking bass; пара 2 — приёмы Capcom-чиптюна   |
| J2  | Mega Drive, светлая        | jump-up     | FM slap-бас, мажор, скорость, jazz-funk-стабы                       |
| J3  | Mega Drive, индустриальная | jump-up     | перегруженный FM-бас, металлическая перкуссия, минор, взрывы        |
| J4  | SNES, геройское с грустью  | jump-up     | лид ведёт парящую мелодию, бас галопом, сэмплерная 16-bit-палитра   |
| J5  | Mega Drive, клубная        | jump-up     | FM-хаус-стабы, ночной город; пара 2 — приправа neuro                |
| J6  | PlayStation, солнечный     | jump-up     | рейв-стабы и ховер начала 90-х, мажор, CD-палитра, SFX бонуса       |
| N1  | jungle 8-bit               | jungle      | amen-брейк, 8-bit-мелодия                                           |
| N2  | jungle с talking bass      | jungle      | amen + jump-up-бас, мост между ядром и jungle                       |
| N3  | роллер                     | dancefloor  | ровный грув, 8-bit-лид как единственная мелодия                     |
| N4 | перелом | halftime | закрыт репризой R3: рифф J1 узнаётся в R3-2a (полный темп, halftime-вставки) и R3-2b (полностью halftime). Какую брать и брать ли — решение отложено |
| N5  | передышка                  | liquid-лайт | без пэдов и пианино, лид ведёт тему, чиптюн-арпеджио вместо пэдов   |
| B1  | шутка                      | breakcore   | 20–30 секунд, бонус-позиция, если войдёт                            |

Если после залпа 1 две оси ядра сольются на слух, одну заменяю на «тёмный» jump-up с half-time-вставкой в середине
трека.

### Раскладка ядра — консоль, настроение, ожидание

Решено 12.09.2026 на отдельной grill-сессии после залпа 1. Правила:

- **База везде jump-up.** Консоль даёт приёмы и ходы, не жанр. В ожидании так и пишем: «jump-up, в который подмешано …».
- **Консоли только четыре:** NES, Mega Drive, SNES, PlayStation. Amiga, Game Boy, Neo Geo, PC Engine, C64 — нет. Вариант «J3 — Amiga» отменён пользователем в пользу Mega Drive.
- **Консоль может повториться, палитра — нет.** Mega Drive стоит трижды (J2, J3, J5), различаем словами: у J2 фанковый slap-бас и мажор, у J3 дисторшн, металл и минор, у J5 хаус-стабы и ровный грув без дисторшна.
- **Названий игр в Style нет** (итог R1). Ориентиры ниже — только для слуха и для строки «Ожидание».
- **Гитары:** гитароподобный перегруженный бас и пауэр-аккорды за басом — можно, гитара-лид и соло — нет. В базовый Exclude добавлено `lead guitar, guitar solo, rock band`, слово `guitars` целиком не пишем, иначе уйдёт и бас.
- **Оценка залпа 2 отвечает на вопрос «попал ли в ожидание»**, а не «хорош ли трек». Для этого под каждым промптом строка «Ожидание».

| Код | Консоль | Настроение | Ожидание: интро и первые 30 с дропа | Темп | Пары |
|---|---|---|---|---|---|
| J1 | NES | наглый, драйв | pulse-лид дублирует talking bass, two-step, аркадные SFX | 175 | 1: эталон J1-1b, Style слово в слово. 2: приёмы Capcom-чиптюна: арпеджио аккордов, имитация slap-баса, драйв |
| J2 | Mega Drive, светлая | скорость, мажор, фанк | FM slap-бас с walking-ходами, яркий FM-лид дублирует хук, jazz-funk-аккорды, ощущение первого уровня платформера | 176 | обе одинаковые, отбор по слуху |
| J3 | Mega Drive, индустриальная | война, безостановочность, антиутопия | перегруженный FM-бас, дисторшн-стабы, металлическая перкуссия, минор, взрывы и выстрелы как SFX | 176 | обе одинаковые |
| J4 | SNES | геройское с грустинкой | pulse-лид ведёт свою парящую мелодию, бас галопом, сэмплерная 16-bit-палитра без оркестра | 172 | 1: чистый лид. 2: плюс guitar-like пауэр-аккорды за басом |
| J5 | Mega Drive, клубная | ночной город, ранние 90-е | FM-хаус-стабы, ровный клубный грув, FM-бас без дисторшна, городские SFX (сирена, стекло) | 172 | 1: без neuro. 2: плюс `evolving modulated texture` |
| J6 | PlayStation | солнечный, скорость, бонус | рейв-стабы и ховер начала 90-х, мажор, CD-палитра вокруг FM-лида, SFX бонуса и чекпойнта | 174 | обе одинаковые |

Ориентиры по слуху (описания жанров — из поиска пользователя, в Style не идут): J1 — Darkwing Duck NES (Capcom-чиптюн, арпеджио аккордов, имитация slap-баса, драйв, хуковая мелодика; нуарную часть не берём, она на отдельный альбом); J2 — Sonic the Hedgehog, Green Hill Zone (16-bit FM, jazz-funk, walking bass, мажор, хуки); J3 — Contra Hard Corps (FM hardcore techno, перегруженный бас, металлическая перкуссия, киберпанк, безостановочность); J4 — Mega Man X (парящая мелодия, галопирующий бас, геройское с меланхолией, пронзительный лид); J5 — Streets of Rage 1 (клубная музыка ранних 90-х, хаус, FM, городская грязца); J6 — Ridge Racer / Ridge Racer Type 4 (дневная трасса, рейв ранних 90-х, мажор, CD-звук).

### Референсы для слуха, не для промптов

Список по памяти, по сети не проверен. Слушать интро и первые 30 секунд каждой темы, как и наши дубли. Композиторов не пишем.

- **NES (J1):** Darkwing Duck, Mega Man 2 (Dr. Wily Stage 1, Air Man), Contra (Jungle), Castlevania (Vampire Killer), Batman NES и Journey to Silius (знаменитый бас, ближе всего к формуле «бас-хук, лид дублирует»), Ninja Gaiden, DuckTales (The Moon).
- **Mega Drive, светлая (J2):** Sonic 1 (Green Hill Zone, Starlight Zone), Sonic 2 (Emerald Hill, Casino Night, Chemical Plant), Sonic 3 (Hydrocity, Ice Cap), Ristar, Rocket Knight Adventures, ToeJam & Earl (чистый FM-фанк), Aero the Acro-bat, Gunstar Heroes.
- **Mega Drive, индустриальная (J3):** Contra Hard Corps, Thunder Force IV, The Revenge of Shinobi и Shinobi III, Streets of Rage 3 (почти хардкор-техно), Batman Mega Drive.
- **Mega Drive, клубная (J5):** Streets of Rage 1 (Fighting in the Street), Streets of Rage 2 (Go Straight, Dreamer, Never Return Alive), Vectorman.
- **SNES (J4):** Mega Man X (Storm Eagle, Spark Mandrill), Donkey Kong Country (Fear Factory, Mine Cart Madness), F-Zero (Mute City, Big Blue), Star Fox (Corneria). Заодно слышно, где на SNES начинаются струнные и духовые, которых нам нельзя.
- **PlayStation (J6):** Ridge Racer, Ridge Racer Type 4, Wipeout (лицензированное техно, палитра эпохи), Jumping Flash!. Отвергнутая альтернатива: Crash Bandicoot / Spyro — мультяшный бонус-уровень, но живая перкуссия воюет с Exclude.
- **Отвергнутые консоли, для памяти:** Amiga (Turrican II, Alien Breed, Project-X, The Chaos Engine) — трекерный звук, был вариант J3; Game Boy (Tetris Type A, Super Mario Land, Kirby's Dream Land) — был вариант J6.

---

## Залп 1 — разведка

### R1. Названия игр в Style

Тот же J1, в конец Style добавлен референс. Сравнивать с лучшим дублем J1. Variety 0.

**Style:**
```
jump up drum and bass, cheeky and hyped, punchy two-step breakbeat, snappy snare, wobbling talking bass riff, big bass hook, 8-bit square lead doubling the bass riff, arcade sound effects, loud club mix, 175 BPM, Contra NES soundtrack style
```

**Lyrics:** как у J1.

**Exclude Styles:** `lyrics, sung words, rap, soft pads, piano`

Оценка: сдвинулся ли звук в сторону военной 8-битной палитры — стаккато, арпеджио, минор. Если ничего не изменилось —
названия игр не работают, в залпе 2 не используем.

### R2. Идиомы эпохи вместо названий

Тот же J1, в Style — описание приёмов NES-музыки. Variety 0.

**Style:**
```
jump up drum and bass, cheeky and hyped, punchy two-step breakbeat, snappy snare, wobbling talking bass riff, 8-bit square lead doubling the bass riff, fast NES-style arpeggiated chord stabs, harmonic minor riff, chromatic run into the drop, short stage clear fanfare before the breakdown, coin and 1-up sounds as fills, loud club mix, 175 BPM
```

**Lyrics:**
```
[Intro: coin and 1-up sounds and drum fill, 8 bars]
[Drop: wobbling bass riff and two-step breakbeat, 8-bit lead doubles the riff, arpeggiated chord stabs]
[Switch: bass riff changes pattern, chromatic run, same drums]
[Fanfare: short stage clear fanfare, then drums out]
[Breakdown: bass riff filtered, snare roll]
[Drop: full hook returns louder]
[Outro: one final bass hit, clean stop]
```

**Exclude Styles:** `lyrics, sung words, rap, soft pads, piano`

Оценка: появились ли арпеджио-стабы и фанфара, стало ли «похоже на игру» сильнее, чем в J1 и R1.

### R3. Cover — реприза в halftime

Под треком «Jump-up 8-bit» в Suno: ⋮ → Remix/Edit → Cover (путь по файлу 02, интерфейс v6 [не проверено]). Max Mode вкл,
Weirdness 35%, Lyrics — оставить как есть.

**Style:**
```
halftime drum and bass, heavy and hypnotic, drum and bass tempo, slow half-time drum pattern at fast tempo, hard snare on the three, deep sub bass, wobbling talking bass stabs, 8-bit square lead keeps the riff, arcade sound effects, loud club mix, 172 BPM
```

**Exclude Styles:** `lyrics, sung words, rap, soft pads, piano, fast rolling drums, full-time breakbeat`

Оценка: узнаётся ли рифф J1. Если да — это N4 и готовый приём против «солянки». Если модель ушла в трэп на 86 BPM —
вторая попытка со Style Influence 90%. Взять или не взять репризу в альбом — решение по слуху и по правовому ощущению.

---

## Залп 1 — ядро jump-up

Все J: v6, Weirdness 45%, Variety Normal, Style Influence 80%, Duration 3:15, по две пары на промпт. Рабочее название в
Suno — код.

### J1. Базовый — контрольная точка

Style — слово в слово как у трека «Jump-up 8-bit» из gamepush (форма Suno). Не менять: это эталон, с которым сравниваем
R1 и R2.

**Style:**
```
Drum and bass, jump-up instrumental at 175 BPM: cheeky hyped energy, 8-bit square lead doubling the bass riff, wobbling talking bass hook, punchy two-step breakbeat, snappy snare, arcade effects, loud club mix with intro, drops, switch, filtered breakdown, and clean outro
```

**Lyrics:**
```
[Intro: arcade sound effects and drum fill, 8 bars]
[Drop: wobbling bass riff and two-step breakbeat, main hook]
[Switch: bass riff changes pattern, same drums]
[Breakdown: drums out, bass riff filtered, snare roll]
[Drop: full hook returns louder]
[Outro: one final bass hit, clean stop]
```

**Exclude Styles:** `lyrics, sung words, rap, soft pads, piano`

Цель двух пар: второй хук того же характера, чтобы в альбоме было два «базовых» на выбор.

### J2. FM-версия — Mega Drive

Палитра 16-бит: FM-синтез, slap bass, металлический лид. Настроение — ночной город, клуб, без названий игр.

**Style:**
```
jump up drum and bass, cheeky and hyped, late night city energy, punchy two-step breakbeat, snappy snare, FM synth slap bass riff, metallic 16-bit FM lead doubling the bass riff, Sega Mega Drive sound chip palette, FM brass stabs, arcade sound effects, loud club mix, 175 BPM
```

**Lyrics:**
```
[Intro: FM brass stab and drum fill, 8 bars]
[Drop: FM slap bass riff and two-step breakbeat, FM lead doubles the riff]
[Switch: bass riff changes pattern, same drums]
[Breakdown: drums out, bass riff filtered, FM brass stabs answer, snare roll]
[Drop: full hook returns louder]
[Outro: one final bass hit, clean stop]
```

**Exclude Styles:** `lyrics, sung words, rap, soft pads, piano, orchestral`

Риск: «Sega Mega Drive sound chip palette» модель может не понять — название консоли как тег [не проверено]. Вторая пара
J2 сразу идёт с заменой этого куска на `16-bit era soundtrack, FM synthesis bassline` (словарь из разведки ниже). Если
обе пары звучат как J1 — FM-палитра через слова не читается, ось J2 заменяем.

### J3. Военный

Марш, стаккато, гармонический минор, взрывы вместо аркадных «монеток». Энергия run-and-gun без названия игры.

**Style:**
```
jump up drum and bass, martial and relentless, punchy two-step breakbeat, snappy snare, machine gun snare rolls, wobbling talking bass riff, 8-bit square lead doubling the bass riff in harmonic minor, staccato NES arpeggio stabs, explosion and gunfire sound effects, loud club mix, 175 BPM
```

**Lyrics:**
```
[Intro: explosion sound effects and machine gun snare roll, 8 bars]
[Drop: wobbling bass riff and two-step breakbeat, 8-bit lead doubles the riff, staccato arpeggio stabs]
[Switch: bass riff changes pattern, same drums, arpeggio stabs faster]
[Breakdown: drums out, bass riff filtered, single explosion, snare roll]
[Drop: full hook returns louder]
[Outro: one final bass hit, clean stop]
```

**Exclude Styles:** `lyrics, sung words, rap, soft pads, piano, happy, major key`

### J4. Аркадный — лид ведёт свою мелодию

Проверка второй формулы хука: лид не дублирует бас, а играет свою запоминающуюся мелодию. Палитра — аркадный автомат:
дудки, взрывы, монетки.

**Style:**
```
jump up drum and bass, hyped and heroic, punchy two-step breakbeat, snappy snare, big wobbling bass hook, 8-bit square lead plays its own catchy melody over the bass, arcade brass stabs, explosion sound effects, coin and jump sound effects as fills, loud club mix, 175 BPM
```

**Lyrics:**
```
[Intro: coin and jump sound effects and drum fill, 8 bars]
[Drop: wobbling bass hook and two-step breakbeat, 8-bit lead plays the main melody on top]
[Switch: bass changes pattern, lead melody continues, same drums]
[Breakdown: drums out, lead melody alone with arcade brass stabs, snare roll]
[Drop: bass hook and lead melody together, louder]
[Outro: one final bass hit, clean stop]
```

**Exclude Styles:** `lyrics, sung words, rap, soft pads, piano`

Оценка: осталась ли мелодия в голове после прослушивания. Если лид опять дублирует бас — формула не читается, у соседей
формулировку усиливаем: `lead melody independent from the bass`.

### J5. С приправой neuro

Бас с меняющимся тембром из neurofunk, но two-step, 8-bit-лид и наглость на месте. Не neurofunk-трек, а jump-up с
холодком.

**Style:**
```
jump up drum and bass, cheeky and menacing, punchy two-step breakbeat, snappy snare, talking bass riff with evolving modulated texture, 8-bit square lead doubling the bass riff, sci-fi laser and alarm sound effects, loud club mix, 175 BPM
```

**Lyrics:**
```
[Intro: alarm sound effect and drum fill, 8 bars]
[Drop: modulated talking bass riff and two-step breakbeat, 8-bit lead doubles the riff]
[Switch: bass texture morphs, pattern changes, same drums]
[Breakdown: drums out, bass texture morphs alone, laser sounds, snare roll]
[Drop: full hook returns louder]
[Outro: one final bass hit, clean stop]
```

**Exclude Styles:** `lyrics, sung words, rap, soft pads, piano, dark ambient, slow`

Риск: модель утащит в neurofunk целиком — барабаны станут синкопированными и «продюсерскими». Если так — убрать
`menacing`, оставить `evolving modulated texture` единственным neuro-словом.

### J6. Солнечный — бонус-уровень

Мажор, прыгучий бас, фанфары «stage clear», монетки. Самый светлый трек ядра.

**Style:**
```
jump up drum and bass, sunny and playful, punchy two-step breakbeat, snappy snare, bouncy wobbling bass riff, bright major key 8-bit square lead doubling the bass riff, short stage clear fanfare, coin and bonus sound effects, loud club mix, 175 BPM
```

**Lyrics:**
```
[Intro: coin sounds and drum fill, 8 bars]
[Drop: bouncy bass riff and two-step breakbeat, bright 8-bit lead doubles the riff]
[Switch: bass riff changes pattern, same drums]
[Fanfare: short stage clear fanfare, drums out]
[Breakdown: bass riff filtered, bonus sound effects, snare roll]
[Drop: full hook returns louder]
[Outro: one final bass hit, clean stop]
```

**Exclude Styles:** `lyrics, sung words, rap, soft pads, piano, dark, minor key`

---

## Итоги залпа 1

Оценки — в журнале ниже. Выводы:

- **Шкала схлопнулась**: 29 четвёрок, одна пятёрка (J1-1b). Причина — у треков не было ожидания. В залпе 2 под каждым промптом строка «Ожидание»: консоль, настроение, что слышно в первые 30 секунд дропа. Оценка отвечает «насколько попал», а не «хорош ли трек». Для мероприятия годятся все, для альбома отбор продолжается.
- **Духовые и струнные лезут везде**: трубы в J2, J4, J6, R2, струнные в J3. Виноваты слова в Style: `FM brass stabs`, `arcade brass stabs`, `stage clear fanfare`, `heroic`, `martial`. Не вкус пользователя. Исправление на все промпты: brass и fanfare из Style убрать, стабы описывать как `FM synth stabs` или `pulse wave stabs`, в Exclude добавить `brass, trumpets, horns, fanfare, orchestral, strings, live instruments`.
- **Голос пролез в R3-1**: в Exclude было `lyrics, sung words, rap`, но не `vocals, singing`. Добавлено во все промпты.
- **Cover-реприза работает**: рифф J1 узнаётся во всех четырёх дублях R3. Style Influence 90% удерживает темп 170+, на 80% модель ушла в 86 BPM и бростеп. N4 с нуля не нужен.
- **J5 без лица**: звучит как J1 с нейрофанком. Нужна своя идентичность через консоль и настроение.
- **J1-1b — эталон ядра**: лид понравился, трек целиком ровный.

### Нуарный пул — на отдельный альбом

R1 и R2 дали кинематографичный 8-битный нуар с озорством, джаз-роком и цирком (по описанию пользователя: ближе к Darkwing Duck и Aero the Acro-bat, чем к Contra). В этот альбом не берём, настроение другое. Ссылки — в журнале, строки R1-a, R1-b, R2-a, R2-b. Рецепт: J1 плюс гармонический минор и NES-арпеджио аккордов; духовые заменить синтетикой.

## Залп 2 — ядро jump-up, второй заход

Промпты по раскладке «консоль + настроение + ожидание» (раздел «Раскладка ядра»). Все J: v6, Weirdness 45%, Variety
Normal, Style Influence 80%, Duration 3:15, по две пары на трек. Рабочее название в Suno — код. Промпты залпа 1 выше
остаются как история, в Suno их больше не запускаем.

Оценка каждого дубля — «попал ли в ожидание» из строки под промптом, плюс обычная цифра 1–5. Где пары различаются
(J1, J4, J5), у второй пары свой Style, помечен.

Базовый Exclude для всех J (таблица «Настройки»):
`lyrics, sung words, rap, vocals, singing, soft pads, piano, brass, trumpets, horns, fanfare, orchestral, strings, live instruments, lead guitar, guitar solo, rock band`

### J1. NES — наглый, драйв

Первая пара — эталон J1-1b, Style слово в слово из sidecar в gamepush, меняется только Exclude. Вторая пара — новый
вариант с приёмами Capcom-чиптюна: арпеджио аккордов, имитация slap-баса, драйв. Нуарных слов нет.

**Ожидание:** NES. Наглый, драйвовый. В интро аркадные SFX и филл, на дропе talking bass и pulse-лид, который его
дублирует нота в ноту, two-step. Во второй паре сверху быстрые арпеджио аккордов, бас с «щелчком» slap-артикуляции.
Ничего живого, никаких духовых.

**Style, пара 1 (эталон, не менять):**
```
Drum and bass, jump-up instrumental at 175 BPM: cheeky hyped energy, 8-bit square lead doubling the bass riff, wobbling talking bass hook, punchy two-step breakbeat, snappy snare, arcade effects, loud club mix with intro, drops, switch, filtered breakdown, and clean outro
```

**Style, пара 2 (Capcom-вариант):**
```
jump up drum and bass, cheeky and hyped, driving action energy, punchy two-step breakbeat, snappy snare, wobbling talking bass riff with 8-bit slap bass articulation, 8-bit square lead doubling the bass riff, fast NES-style arpeggiated chord stabs, catchy hook-heavy riff, arcade sound effects, loud club mix, 175 BPM
```

**Lyrics, обе пары:**
```
[Intro: arcade sound effects and drum fill, 8 bars]
[Drop: wobbling bass riff and two-step breakbeat, 8-bit lead doubles the riff, main hook]
[Switch: bass riff changes pattern, same drums]
[Breakdown: drums out, bass riff filtered, snare roll]
[Drop: full hook returns louder]
[Outro: one final bass hit, clean stop]
```

Во второй паре в тег `[Drop]` дописать `, arpeggiated chord stabs on top`.

**Exclude Styles:** базовый.

**Пара 3 — структура B.** Style и Exclude пары 2 (Capcom-вариант; эталонный Style пары 1 не годится, в нём форма
зашита словами). Схема «босс-раш»: без раскачки, дроп сразу, ложная концовка со звуком 1-up посреди трека, финальный
дроп с контр-риффом лида.

**Ожидание, пара 3:** палитра пары 2. Барабанный филл и сразу дроп; после switch всё обрывается на одном ударе баса,
звук 1-up, пауза, и хук возвращается громче; в последнем дропе лид играет не рифф баса, а свой ответ ему.

**Lyrics, пара 3 (структура B):**
```
[Intro: drum fill only, 4 bars]
[Drop: wobbling bass riff and two-step breakbeat immediately, 8-bit lead doubles the riff, 32 bars]
[Switch: bass riff changes pattern, arpeggiated chord stabs enter, 16 bars]
[Fake ending: everything stops on one bass hit, 1-up sound effect, 2 bars of silence]
[Drop: full hook returns louder, 32 bars]
[Breakdown: drums out, bass riff filtered, snare roll, 8 bars]
[Drop: final hook, 8-bit lead plays a counter-riff answering the bass, 16 bars]
[Outro: one final bass hit, clean stop]
```

### J2. Mega Drive, светлая сторона — скорость, мажор, фанк

**Ожидание:** Mega Drive. Первый уровень платформера: солнечно, быстро, мажор. На дропе FM slap-бас с walking-ходами
между фразами, яркий металлический FM-лид дублирует хук, между ними jazz-funk-аккорды FM-стабами. SFX — монетки и
прыжки. Ни труб, ни струнных: стабы синтетические.

**Style:**
```
jump up drum and bass, bright and fast, speedy platformer energy, bright major key, punchy two-step breakbeat, snappy snare, FM synth slap bass riff with walking bass runs, bright metallic FM lead doubling the bass riff, jazz-funk FM synth chord stabs, 16-bit era FM synthesis palette, coin and jump sound effects, loud club mix, 176 BPM
```

**Lyrics:**
```
[Intro: FM chord stab and coin sounds, drum fill, 8 bars]
[Drop: FM slap bass riff and two-step breakbeat, FM lead doubles the riff, jazz-funk chord stabs]
[Switch: bass riff changes pattern, walking bass run, same drums]
[Breakdown: drums out, bass riff filtered, FM chord stabs answer, snare roll]
[Drop: full hook returns louder]
[Outro: one final bass hit, clean stop]
```

**Exclude Styles:** базовый плюс `minor key, dark`.

**Пара 3 — без slap-баса и jazz-funk.** По первым дублям пары 1: настроение классное, но slap-бас и jazz-funk-стабы
делают музыку слишком сумасшедшей. Бас возвращается к обычному talking bass, стабы простые и яркие. Sonic остаётся
через мажор, скорость и хуковый FM-лид.

**Ожидание, пара 3:** то же, что у пар 1–2, но бас ровный, без щелчков и walking-ходов, аккорды простые, без джазовых
надстроек. Спокойнее по гармонии, такая же быстрая.

**Style, пара 3:**
```
jump up drum and bass, bright and fast, speedy platformer energy, bright major key, punchy two-step breakbeat, snappy snare, wobbling FM synth talking bass riff, bright metallic FM lead doubling the bass riff with a catchy hook, simple bright FM synth chord stabs, 16-bit era FM synthesis palette, coin and jump sound effects, loud club mix, 176 BPM
```

**Lyrics, пара 3:**
```
[Intro: FM chord stab and coin sounds, drum fill, 8 bars]
[Drop: wobbling FM bass riff and two-step breakbeat, FM lead doubles the riff, bright chord stabs]
[Switch: bass riff changes pattern, same drums]
[Breakdown: drums out, bass riff filtered, FM chord stabs answer, snare roll]
[Drop: full hook returns louder]
[Outro: one final bass hit, clean stop]
```

**Exclude Styles, пара 3:** базовый плюс `minor key, dark, slap bass, jazz, funk`.

**Пара 4 — структура B.** Все J залпа 2 идут по одному скелету (intro → drop → switch → breakdown → drop → outro,
схема A, унаследована от эталона J1). Здесь первая проба другой формы: Style и Exclude — как у пары 3, меняется только
Lyrics. Лид заявляет хук до баса, как мелодичное интро платформера; вместо switch — второй дроп со вторым хуком; outro
лидом соло. Счёт тактов вне интро — эксперимент, читает ли его v6 [не проверено, файл 02].

**Ожидание, пара 4:** палитра пары 3. Первые секунды — FM-лид один с монетками, без барабанов; барабаны входят на build;
на первом дропе бас берёт хук лида; второй дроп даёт новый паттерн баса и второй хук; в конце лид один раз играет хук
без ритма и стоп.

**Lyrics, пара 4 (структура B):**
```
[Intro: FM lead plays the main hook alone over coin sounds, no drums, 8 bars]
[Build: two-step drums enter, bass hinted, snare roll, 8 bars]
[Drop: wobbling FM bass riff and two-step breakbeat, FM lead doubles the riff, 32 bars]
[Second drop: bass riff changes to a new pattern, FM lead answers with a second hook, 16 bars]
[Breakdown: drums out, bright chord stabs and lead hook only, jump sound effect, 8 bars]
[Drop: first riff returns louder, both hooks trade, 32 bars]
[Outro: drums stop dead, FM lead plays the hook once alone, clean stop]
```

Оценка пары 4 против пары 3: слышна ли разница формы, держит ли модель второй хук, не потерялся ли jump-up на длинном
дропе. Свои схемы B у J1, J3, J4, J5, J6 — под каждым промптом, пара 3.

### J3. Mega Drive, индустриальная сторона — война, безостановочность

**Ожидание:** Mega Drive. Антиутопия, run-and-gun без передышки, минор. На дропе перегруженный FM-бас, звучащий почти как
гитара, дисторшн-лид дублирует рифф, металлические перкуссионные удары поверх two-step. SFX — взрывы и выстрелы, сирена
в интро. Никаких струнных: минор держится синтетикой.

**Style:**
```
jump up drum and bass, relentless and dystopian, gritty industrial 16-bit FM synthesis, punchy two-step breakbeat, snappy snare, metallic percussion hits, overdriven FM bass riff, distorted FM lead doubling the bass riff in minor key, aggressive distorted FM synth stabs, explosion and gunfire sound effects, loud club mix, 176 BPM
```

**Lyrics:**
```
[Intro: alarm siren and explosion, machine gun snare roll, 8 bars]
[Drop: overdriven bass riff and two-step breakbeat, distorted lead doubles the riff, metallic hits]
[Switch: bass riff changes pattern, same drums, distorted stabs faster]
[Breakdown: drums out, overdriven bass alone, metallic hits, single explosion, snare roll]
[Drop: full hook returns louder]
[Outro: one final bass hit, clean stop]
```

**Exclude Styles:** базовый плюс `happy, major key`.

**Пара 3 — структура B.** Style и Exclude пар 1–2. Схема «без передышки»: breakdown нет вовсе, вместо него восемь
тактов half-time с перегруженным басом на третьей доле, потом two-step возвращается на полной скорости. Заканчивается
взрывом, а не чистым ударом.

**Ожидание, пара 3:** палитра та же. Сирена, четыре такта снейр-ролла и сразу дроп; после switch барабаны падают в
половинный темп, бас бьёт тяжёлыми стабами; возврат в two-step должен ощущаться как рывок; в конце всё обрывается
взрывом.

**Lyrics, пара 3 (структура B):**
```
[Intro: alarm siren, machine gun snare roll, 4 bars]
[Drop: overdriven bass riff and two-step breakbeat, distorted lead doubles the riff, metallic hits, 32 bars]
[Switch: bass riff changes pattern, distorted stabs faster, 16 bars]
[Half-time: drums drop to half-time, overdriven bass stabs on the three, 8 bars]
[Drop: two-step returns at full speed, full hook louder, 32 bars]
[Switch: new bass pattern, metallic hits doubled, 16 bars]
[Outro: everything cuts on a single explosion, end]
```

### J4. SNES — геройское с грустинкой, лид ведёт мелодию

Вторая формула хука: лид не дублирует бас, а играет свою мелодию. Первая пара — чистый pulse-лид. Вторая — тот же лид
плюс сэмплерные гитароподобные пауэр-аккорды за басом (ход в духе Metrik, но в консольной палитре).

**Ожидание:** SNES. Геройское с меланхолией, минор с просветами. В интро чиптюн-арпеджио и лид один раз играет мелодию.
На дропе бас галопом, пронзительный pulse-лид ведёт парящую мелодию сверху, палитра сэмплерная, 16-bit, без оркестра.
Во второй паре под лидом слышны короткие пауэр-аккорды, повторяющие бас. Гитары-соло нет.

**Style, пара 1:**
```
jump up drum and bass, hyped and bittersweet, soaring melodic energy, punchy two-step breakbeat, snappy snare, galloping wobbling bass riff, piercing pulse wave lead plays its own soaring melody over the bass, sampled 16-bit console synth palette, bright chiptune arpeggios, jump and charged shot sound effects, loud club mix, 172 BPM
```

**Style, пара 2 (плюс пауэр-аккорды):**
```
jump up drum and bass, hyped and bittersweet, soaring melodic energy, punchy two-step breakbeat, snappy snare, galloping wobbling bass riff, sampled 16-bit guitar-like power chord stabs following the bass, piercing pulse wave lead plays its own soaring melody over the bass, sampled 16-bit console synth palette, bright chiptune arpeggios, jump and charged shot sound effects, loud club mix, 172 BPM
```

**Lyrics, обе пары:**
```
[Intro: chiptune arpeggio alone, pulse lead plays the melody once, drum fill]
[Drop: galloping bass riff and two-step breakbeat, pulse lead plays the main melody on top]
[Switch: bass changes pattern, lead melody continues, same drums]
[Breakdown: drums out, lead melody alone with arpeggios, snare roll]
[Drop: bass riff and lead melody together, louder]
[Outro: one final bass hit, clean stop]
```

Во второй паре в тег `[Drop]` дописать `, power chord stabs follow the bass`.

**Exclude Styles:** базовый.

Оценка: осталась ли мелодия в голове. Если лид опять дублирует бас — усиливаем `lead melody independent from the bass`.

**Пара 3 — структура B.** Style и Exclude пары 1 (чистый лид). Схема «тема прежде всего»: длинное интро, лид играет
всю мелодию до барабанов и баса, switch нет, в breakdown лид играет вторую мелодию, финальный дроп возвращает первую.
Outro — концовка мелодии соло над арпеджио.

**Ожидание, пара 3:** палитра та же. Первые 16 тактов без барабанов: арпеджио, затем мелодия целиком; барабаны входят
отдельной секцией; на дропе бас галопом под уже знакомой мелодией; в breakdown другая мелодия, и по ней слышно, что
лид ведёт, а не дублирует.

**Lyrics, пара 3 (структура B):**
```
[Intro: chiptune arpeggio alone, no drums, 8 bars]
[Intro: pulse lead plays the full melody over the arpeggio, no drums, 8 bars]
[Build: two-step drums enter, snare roll, 8 bars]
[Drop: galloping bass riff and two-step breakbeat, pulse lead plays the main melody on top, 32 bars]
[Breakdown: drums out, pulse lead plays a second melody with arpeggios, 16 bars]
[Drop: bass riff and first melody together, louder, 32 bars]
[Outro: drums stop, pulse lead plays the melody ending alone over the arpeggio, clean stop]
```

**Пара 4 — 80/20 от пары 2.** По оценке пользователя после первых дублей: J4 «вообще крутая», особенно часть с
пауэр-аккордами, и главное — геройское с меланхолией, минор с просветами, мелодия как соло. Это то, что у него
ассоциируется с играми. Задача пары 4: сохранить 80% пары 2 и добавить 20% вариативности. Style пары 2 плюс два
уточнения: `melancholic minor melody with brief major key lifts` и `expressive pulse wave lead solo`. Структура A, но
перед последним дропом секция соло: барабаны и бас идут, лид играет развёрнутую меланхоличную мелодию с подъёмом в
мажор. Exclude тот же, `guitar solo` в базовом списке страхует от гитары на слове solo.

**Ожидание, пара 4:** палитра пары 2, пауэр-аккорды за басом на месте. Первые 30 секунд дропа — как у пары 2. Новое
слышно после второго дропа: лид уходит в соло на несколько фраз, мелодия минорная с просветом в мажор, потом главная
мелодия возвращается на один последний дроп.

**Style, пара 4:**
```
jump up drum and bass, hyped and bittersweet, soaring melodic energy, melancholic minor melody with brief major key lifts, punchy two-step breakbeat, snappy snare, galloping wobbling bass riff, sampled 16-bit guitar-like power chord stabs following the bass, piercing pulse wave lead plays its own soaring melody over the bass, expressive pulse wave lead solo, sampled 16-bit console synth palette, bright chiptune arpeggios, jump and charged shot sound effects, loud club mix, 172 BPM
```

**Lyrics, пара 4:**
```
[Intro: chiptune arpeggio alone, pulse lead plays the melody once, drum fill]
[Drop: galloping bass riff and two-step breakbeat, power chord stabs follow the bass, pulse lead plays the main melody on top]
[Switch: bass changes pattern, lead melody continues, same drums]
[Breakdown: drums out, lead melody alone with arpeggios, snare roll]
[Drop: bass riff, power chords and lead melody together, louder]
[Solo: drums and bass continue, pulse lead plays an expressive melancholic solo, melody rises into a brief major key lift]
[Drop: main melody returns one last time with power chords]
[Outro: one final bass hit, clean stop]
```

**Exclude Styles, пара 4:** базовый.

### J5. Mega Drive, клубная сторона — ночной город, ранние 90-е

Первая пара — без neuro. Вторая — единственное neuro-слово `evolving modulated texture` у баса. Сравниваем на одном
настроении: приправа либо остаётся в альбоме, либо уходит.

**Ожидание:** Mega Drive. Ночной город, клуб ранних 90-х, дерзко, но без дисторшна. В интро сирена и FM-хаус-стабы. На
дропе ровный клубный грув, FM-бас ведёт рифф, металлический FM-лид его дублирует, хаус-стабы отвечают между фразами.
Барабаны остаются two-step, не four-on-the-floor. Во второй паре бас меняет тембр по ходу риффа.

**Style, пара 1:**
```
jump up drum and bass, cheeky and streetwise, late night city club energy, early 90s house groove, punchy two-step breakbeat, snappy snare, driving FM synth bass riff, metallic FM lead doubling the bass riff, FM house chord stabs, 16-bit era FM synthesis palette, siren and breaking glass sound effects, loud club mix, 172 BPM
```

**Style, пара 2 (плюс neuro):**
```
jump up drum and bass, cheeky and streetwise, late night city club energy, early 90s house groove, punchy two-step breakbeat, snappy snare, driving FM synth bass riff with evolving modulated texture, metallic FM lead doubling the bass riff, FM house chord stabs, 16-bit era FM synthesis palette, siren and breaking glass sound effects, loud club mix, 172 BPM
```

**Lyrics, обе пары:**
```
[Intro: siren and FM house chord stabs, drum fill, 8 bars]
[Drop: FM bass riff and two-step breakbeat, FM lead doubles the riff, house stabs answer]
[Switch: bass riff changes pattern, same drums]
[Breakdown: drums out, bass riff filtered, house stabs, breaking glass, snare roll]
[Drop: full hook returns louder]
[Outro: one final bass hit, clean stop]
```

Во второй паре в тег `[Switch]` дописать `, bass texture morphs`.

**Exclude Styles:** базовый плюс `four on the floor, house tempo, dark ambient, slow`.

Риск: `house groove` утащит барабаны в 4/4. Если так в обеих парах — убрать `early 90s house groove`, оставить
`FM house chord stabs` единственным словом про хаус.

**Пара 3 — структура B.** Style и Exclude пары 1 (без neuro). Схема «диджейская»: длинное интро на одних барабанах,
потом стабы, потом бас; в середине секция без баса, где грув держат стабы и барабаны; outro — барабаны одни восемь
тактов и стоп, как под сведение.

**Ожидание, пара 3:** палитра та же. Первые 8 тактов только two-step и сирена, потом хаус-стабы, и лишь на дропе бас;
посреди трека бас уходит, а грув не проваливается; в конце бас выключается раньше барабанов.

**Lyrics, пара 3 (структура B):**
```
[Intro: two-step drums alone with siren, 8 bars]
[Intro: FM house chord stabs join the drums, bass hinted, 8 bars]
[Drop: FM bass riff and two-step breakbeat, FM lead doubles the riff, house stabs answer, 32 bars]
[Stabs section: bass out, drums keep rolling, house chord stabs carry the groove, 16 bars]
[Drop: bass riff returns with a new pattern, FM lead doubles, louder, 32 bars]
[Breakdown: drums out, bass riff filtered, breaking glass, snare roll, 8 bars]
[Drop: full hook, 16 bars]
[Outro: bass out, drums roll alone, 8 bars, stop dead]
```

### J6. PlayStation — солнечный, скорость, бонус

**Ожидание:** PlayStation. Дневная трасса, бонус-этап, мажор, эйфория без пэдов. В интро звук чекпойнта и рейв-стаб. На
дропе прыгучий бас, яркий FM-лид дублирует рифф, вокруг рейв-стабы и ховер начала 90-х, палитра чистая, CD-качества,
не чиповая. SFX — бонусы и чекпойнт. Самый светлый и самый «не 8-битный» трек ядра.

**Style:**
```
jump up drum and bass, sunny and euphoric, daytime racing bonus stage energy, bright major key, punchy two-step breakbeat, snappy snare, bouncy wobbling bass riff, bright FM lead doubling the bass riff, early 90s rave synth stabs, hoover synth riff, clean 32-bit CD-quality sampled synth palette, checkpoint and bonus sound effects, loud club mix, 174 BPM
```

**Lyrics:**
```
[Intro: checkpoint sound effect and rave synth stab, drum fill, 8 bars]
[Drop: bouncy bass riff and two-step breakbeat, FM lead doubles the riff, rave stabs answer]
[Switch: bass riff changes pattern, hoover riff, same drums]
[Breakdown: drums out, hoover riff alone, bonus sound effects, snare roll]
[Drop: full hook returns louder]
[Outro: one final bass hit, clean stop]
```

**Exclude Styles:** базовый плюс `trance, dark, minor key`.

Риск: `euphoric` может позвать трансовые пэды. `soft pads` уже в Exclude; если пэды всё равно пришли — заменить
`euphoric` на `playful`.

**Пара 3 — структура B.** Style и Exclude пар 1–2. Схема «гонка»: два дропа подряд без switch (второй ведёт ховер),
звук чекпойнта на границах секций, в конце «финальный круг»: бас удваивает скорость, хэты удваиваются.

**Ожидание, пара 3:** палитра та же. Чекпойнт, рейв-стаб нарастает со снейр-роллом; первый дроп короткий, сразу
второй с ховером в роли хука; после breakdown длинный полный дроп; перед концом чекпойнт и заметное ускорение
внутри того же темпа.

**Lyrics, пара 3 (структура B):**
```
[Intro: checkpoint sound effect, rave synth stab builds with snare roll, no bass, 8 bars]
[Drop: bouncy bass riff and two-step breakbeat, FM lead doubles the riff, 16 bars]
[Second drop: hoover riff takes over the hook, bass follows it, 16 bars]
[Breakdown: drums out, rave stabs and bonus sound effects, snare roll, 8 bars]
[Drop: full hook returns louder, rave stabs and hoover trade, 32 bars]
[Final lap: checkpoint sound effect, bass riff doubles its speed, hi-hats doubled, 16 bars]
[Outro: one final bass hit, bonus sound effect, clean stop]
```

---

## Залп 2 — соседи (черновик, правлю после оценок залпа 1)

Все N: v6, Weirdness 35%, Variety Normal, Style Influence 80%, Duration 3:45. Инварианты на месте: лид, SFX, loud club
mix.

### N1. Jungle 8-bit

Файл 04 §2 плюс 8-bit-лид, который ведёт мелодию — вторая формула хука.

**Style:**
```
oldschool jungle, raw and hyped, chopped amen break, fast rolling drums, deep reggae sub bass, rave stabs, 8-bit square lead carries the main melody, siren samples, arcade sound effects, dusty sampler sound, loud club mix, 170 BPM
```

**Lyrics:**
```
[Intro: rave stab and siren, then amen break cuts in immediately]
[Drop: chopped amen break and reggae sub bass, 8-bit lead plays the main melody]
[Breakdown: drums drop out, 8-bit lead alone with siren]
[Drop: amen break returns with double-time chops, bass heavier, lead melody on top]
[Outro: drums stop dead, one siren, end]
```

**Exclude Styles:** `lyrics, sung verses, rap verses, vocals, singing, timestretched vocal shouts, clean modern mix, trap, brass, trumpets, horns, fanfare, orchestral, strings, live instruments, lead guitar, guitar solo, rock band`

Заметка: `timestretched vocal shouts` из файла 04 убраны в Exclude — на альбоме без вокала выкрики будут торчать. Если
jungle без них звучит пусто — вернуть в Style.

Заметка 12.09.2026: `air horn` убран из Style и Lyrics — слово `horn` спорило с `horns` в Exclude. Наблюдение
пользователя по первым дублям: мелодия у N1 та же, что у мажорных J, хотя здесь лид ведёт свою мелодию, а не дублирует
бас. Значит, формула «лид дублирует бас» — не главная причина схождения мелодий; остаются словарь промпта, мажор в
чиптюне. Personalize / My Taste проверен — выключен, эта причина отпадает.

**Пара 2 — против схождения мелодии.** Первая проба рецепта: лад, характер мелодии словами, Weirdness выше. Мажор
уходит, jungle и так минорный по природе; мелодии заданы синкопы, широкие скачки и перекличка со стабами; в Exclude
`major key` и `simple pentatonic melody`, чтобы отрезать заезженный контур. Weirdness 55% вместо 35%, остальное как у
всех N. Если мелодия ушла от мажорных J — тот же рецепт на J2 и J6.

Заодно тест тональности: вместо `minor key` в Style стоит `in E minor`. Гайды сообщества советуют писать тональность в
Style, но держит ли её v6 — [не проверено]. Ми минор выбран как типичная тональность jungle-сабов. Проверка — тюнер или
анализ тональности по готовому треку. Если держит — у каждого трека альбома своя тональность, и мелодиям труднее
совпасть.

**Ожидание, пара 2:** amen, reggae sub, rave-стабы как в паре 1, но лид минорный, рваный, со скачками через октаву, и
он отвечает стабам, а не поёт поверх ровной пентатоникой. Первые 30 секунд дропа: мелодия не должна напоминать J2 и J6.
Тональность — ми минор, проверить по тюнеру.

**Style, пара 2:**
```
oldschool jungle, raw and dark, in E minor, chopped amen break, fast rolling drums, deep reggae sub bass, rave stabs, syncopated 8-bit square lead melody with wide interval leaps, call and response between the lead and the rave stabs, siren samples, arcade sound effects, dusty sampler sound, loud club mix, 170 BPM
```

**Lyrics, пара 2:**
```
[Intro: rave stab and siren, then amen break cuts in immediately]
[Drop: chopped amen break and reggae sub bass, syncopated 8-bit lead melody answers the rave stabs]
[Breakdown: drums drop out, 8-bit lead alone, wide leaps, siren]
[Drop: amen break returns with double-time chops, bass heavier, lead melody on top]
[Outro: drums stop dead, one siren, end]
```

**Exclude Styles, пара 2:** как у пары 1 плюс `major key, happy, simple pentatonic melody`.

### N2. Jungle с talking bass — мост к ядру

**Style:**
```
oldschool jungle, raw and cheeky, in G minor, chopped amen break, fast rolling drums, wobbling talking bass riff, 8-bit square lead doubling the bass riff, rave stabs, siren samples, arcade sound effects, loud club mix, 172 BPM
```

**Lyrics:**
```
[Intro: siren and amen break fill, 8 bars]
[Drop: chopped amen break and wobbling talking bass riff, 8-bit lead doubles the riff]
[Switch: bass riff changes pattern, amen chops double-time]
[Breakdown: drums out, bass riff filtered, rave stab, snare roll]
[Drop: full hook returns louder]
[Outro: drums stop dead, one bass hit, end]
```

**Exclude Styles:** `lyrics, sung words, rap, vocals, singing, soft pads, piano, clean modern mix, brass, trumpets, horns, fanfare, orchestral, strings, live instruments, lead guitar, guitar solo, rock band`

Заметка 12.09.2026: в Style добавлено `in G minor` — вторая проверка тональности после N1 (там ми минор). Разные
тональности у соседних треков, чтобы мелодиям было труднее совпасть. Держит ли v6 тональность — [не проверено],
смотреть по тюнеру.

### N3. Роллер

Файл 04 §4 с 8-bit-лидом как единственной мелодией.

**Style:**
```
dancefloor drum and bass roller, driving and steady, in F minor, tight punchy breakbeat, sharp snare, deep rolling sub bass, minimal stabs, 8-bit square lead plays a short repeating motif, arcade sound effects sparse, loud club mix, relentless groove, 174 BPM
```

**Lyrics:**
```
[Intro: filtered breakbeat, bass hinted, 16 bars]
[Drop: full breakbeat and rolling sub bass, 8-bit motif enters]
[Roll: same groove continues, subtle variation in hats, motif repeats]
[Breakdown: bass out, drums keep rolling, arcade sound effect]
[Drop: bass returns, full groove to the end]
[Outro: drums cut to a single hit, end]
```

**Exclude Styles:** `lyrics, sung words, rap, vocals, singing, melody-heavy, soft pads, piano, brass, trumpets, horns, fanfare, orchestral, strings, live instruments, lead guitar, guitar solo, rock band`

Заметка 12.09.2026: в Style добавлено `in F minor` — тест тональности, как у N1 (ми минор) и N2 (соль минор).
[не проверено], смотреть по тюнеру.

### N4. Halftime — перелом

**Закрыт после залпа 1:** R3 сработал, позицию занимает реприза (R3-2a или R3-2b, решение отложено). Промпт ниже оставлен на случай, если реприза в альбом не пойдёт по правовому ощущению. При перегенерации R3 — Style Influence 90%, в Exclude добавить `vocals, singing`.

**Style:**
```
halftime drum and bass, heavy and hypnotic, drum and bass tempo, slow half-time drum pattern at fast tempo, hard snare on the three, deep sub bass, wobbling talking bass stabs, 8-bit square lead plays a slow motif, arcade sound effects, loud club mix, 172 BPM
```

**Lyrics:**
```
[Intro: sub bass and single hi-hat pattern, no snare, arcade sound effect]
[Drop: half-time drums, deep sub and wobbling bass stabs, 8-bit motif]
[Variation: extra percussion, bass pattern changes]
[Break: drums out, sub drone, 8-bit motif alone]
[Drop: half-time beat returns, heavier]
[Outro: drums stop, sub fades]
```

**Exclude Styles:** `lyrics, sung words, rap, vocals, singing, soft pads, piano, fast rolling drums, full-time breakbeat, trap, brass, trumpets, horns, fanfare, orchestral, strings, live instruments, lead guitar, guitar solo, rock band`

### N5. Liquid-лайт — передышка

Liquid без пэдов и пианино: их место занимают чиптюн-арпеджио, лид ведёт тему.

**Style:**
```
liquid drum and bass, warm and bright, in D major, smooth rolling breakbeat, crisp snare, deep round sub bass, 8-bit square lead carries a melodic theme, plucked chiptune arpeggios instead of pads, sparse arcade sound effects, clean polished mix, 174 BPM
```

**Lyrics:**
```
[Intro: chiptune arpeggio alone, no drums, lead plays the theme once]
[Build: hi-hats and soft snare roll enter]
[Drop: full rolling breakbeat and sub bass, lead theme on top]
[Breakdown: drums out, arpeggio and lead theme]
[Drop: full beat returns, arpeggio and lead together]
[Outro: drums fade, arpeggio alone, soft end]
```

**Exclude Styles:** `lyrics, sung words, rap, vocals, singing, piano, string pads, aggressive bass, distortion, brass, trumpets, horns, fanfare, orchestral, strings, live instruments, lead guitar, guitar solo, rock band`

Заметка: `clean polished mix` вместо `loud club mix` — единственное отступление от инварианта, иначе передышки не будет.
Если трек выпадает из альбома по громкости — вернуть `loud club mix`.

Заметка 12.09.2026: в Style добавлено `in D major` — тест тональности, единственный мажорный сосед. Если мелодия
совпадёт с мажорными J даже в другой тональности — схождение не про тональность, а про мажорный чиптюн как таковой.
[не проверено], смотреть по тюнеру.

### B1. Breakcore-вставка

Файл 04 §11 как есть. Duration 0:25, Weirdness 65%.

**Style:**
```
drill and bass breakcore burst, chaotic and playful, hyper-chopped amen break, stuttering glitch edits, rubbery acid bass, 8-bit melody fragment, arcade sound effects, no intro build-up, single short burst, abrupt stop, 180 BPM
```

**Lyrics:**
```
[Burst: chopped amen and acid bass explode immediately, 8-bit melody fragment]
[End: abrupt stop on a glitch]
```

**Exclude Styles:** `vocals, singing, rap, lyrics, slow intro, ambient, steady groove, brass, trumpets, horns, fanfare, orchestral, strings, live instruments, lead guitar, guitar solo, rock band`

---

## Разведка по сети

Два вопроса, лимит 30 минут. Источники — гайды сообщества и справочники промптов; собственные страницы стилей suno.com
отдаются пустыми без JavaScript, оттуда ничего не взял.

**1. Читает ли Suno названия игр в Style.** Прямого ответа нигде нет — ни подтверждения, ни опровержения. Что есть:
- Имена артистов Suno блокирует, и по сообщениям 2026 года за них прилетают страйки за имперсонацию [сообщ]. Это
  подтверждает решение не писать композиторов.
- Гайды по игровой музыке названия игр не используют вообще. Формулировка одного справочника: теги «describe format,
  era, and scene function so you can build an original direction without naming a protected score» [сообщ]. То есть
  сообщество обходится без названий — либо потому что не работают, либо из осторожности.
- Вывод: R1 остаётся единственной проверкой. Если сдвига нет — названия игр закрываем окончательно.

**2. Как описывают FM-звук Mega Drive.** Промптов именно под Suno с YM2612 или Mega Drive нет. Рабочий словарь, который
встречается в гайдах:
- `16-bit era soundtrack`, `16-bit`, `retro game music`, `8-bit music`, `chiptune` — теги из справочника
  usesuno [сообщ].
- «16-bit» в промпте, по одному из гайдов, «triggers the warmer, FM synthesis sound of 90s
  consoles» [сообщ, один источник, не проверено].
- Описание звука в стиле карточек Suno: `FM synthesis basslines`, `square wave and sawtooth oscillators with vibrato`,
  `percussion typical of 90s console sound chips` [сообщ].
- Пример 16-bit-промпта из гайда: `16-bit SNES soundtrack, retro synths, nostalgic, platformer music, 120 BPM` [сообщ].
  SNES — сэмплерный чип, не FM; для Mega Drive словосочетание не годится.

Что меняю по итогам: во второй паре J2 вместо `Sega Mega Drive sound chip palette` ставлю
`16-bit era soundtrack, FM synthesis bassline` — это словарь, который сообщество считает рабочим, а название
консоли — [не проверено].

Источники: [usesuno — game music list](https://usesuno.com/guide/reference/game-music-list), [hookgenius — gaming prompts](https://hookgenius.app/learn/suno-gaming-music-prompts/), [hookgenius — Suno for game developers](https://hookgenius.app/learn/suno-for-game-developers/), [sunoprompt — gaming genre](https://sunoprompt.com/music-style-genre/gaming-music-genre), [roo.beehiiv — artist style prompts](https://roo.beehiiv.com/p/suno-artist-style-prompts), [songsmith — reference without naming the artist](https://songsmith.studio/blog/suno-reference-track-prompt-builder).

## Порядок треков — принцип

Энергетическая дуга: открываем сильнейшим J, соседи-передышки на позициях 4 и 8, halftime ближе к концу как перелом, B1
перед финальным треком, закрываем вторым по силе J. Итоговый порядок — после отбора, правится на слух.

## Названия — принцип

После отбора. По палитре и настроению, можно на английском, под игровой жаргон без конкретики: Insert Coin, Stage Clear,
Continue?. Имён игр и персонажей нет.

## Журнал

### Как оценивать

Шкала — одна цифра, без долгих раздумий:

| Оценка | Смысл                                                      |
|--------|------------------------------------------------------------|
| 5      | точно в альбом, ничего менять не надо                      |
| 4      | в альбом, если не появится лучше                           |
| 3      | что-то есть, но промпт надо править — напишите, что именно |
| 2      | не то, но понятно, почему                                  |
| 1      | мусор, дальше не слушал                                    |

«Что не так» — одно-два слова из списка, можно несколько через запятую: `бас`, `лид`, `барабаны`, `SFX`, `структура`,
`скучно`, `вокал прорвался`, `не тот жанр`, `слишком похоже на J1`. Свои слова тоже годятся.

Ссылка — share-ссылка из Suno, необязательно. Слушать целиком не нужно: интро и первые 30 секунд дропа решают, дальше —
только если 4–5.

Первый дубль пары — `a`, второй — `b`. У J два запуска: `J1-1a`, `J1-1b` — первая пара, `J1-2a`, `J1-2b` — вторая.

### Залп 1 — 12.09.2026

| Код   | Оценка | Что не так | Ссылка                              | Примечания                                                                                                                                                                                                                                                                      |
|-------|--------|------------|-------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| R1-a  | 4      |            | https://suno.com/s/5CO2rlAPP26bmcHd | Точно повлияло, но Contra тут не пахнет, больше на Darkwing duck NES похожа музыка, скорее совпадение. Но справедливости ради, музыка стала интересней. Как мне google описал "кинематографичный 8-битный нуар-синтвейв с элементами джаз-рока"                                 |
| R1-b  | 4      |            | https://suno.com/s/cCRQIiB6E7dMHlvj | Тут так же conra не пахнет.                                                                                                                                                                                                                                                     |
| R2-a  | 4      |            | https://suno.com/s/FlE7GqIYXBCW0HnG | Это достойно, но всё равно 4 из за труб. Это больше похоже на музыку из игр, что-то вроде Sega Aero The Acro-bat. Тоже нуарность присутствует, это надо в отдельный альбом выделять. Но аркестровость я бы попробовал заменить на синтетику, кажется живые трубы вообще не моё. |
| R2-b  | 4      |            | https://suno.com/s/kQJpTP0klXD4lkG4 | Тоже, есть нуарное настроение, озорное, не давящие, цырковое такое, (Sega Aero The Acro-bat). Трубы минус.                                                                                                                                                                      |
| R3-1a | 4      |            | https://suno.com/s/Ol6Mt9HKpb6mvydB | 86 BPM в целом ремикс удался, стало похоже на бростеп, прикольно, референсный трек узнаётся, откуда-то вылез женский голос, но вроде прикольно                                                                                                                                  |
| R3-1b | 4      |            | https://suno.com/s/y8XiToohaW1X8awM | этот уже не стал весь трек до 86 BPM опускать, есть местами отдельные вставки, оригинальный трек узнаётся, тоже считаю что ремикс удался. Тоже женски голос откуда-то влез, голос дублирует/заменяет лид                                                                        |
| R3-2a | 4      |            | https://suno.com/s/52snhKURG9PjR0ks | Style Influence 90%. Темп полноценный Тут оригинальные вставки есть. Т.е оригинальный риф узнаётся и есть халфтайм вставки                                                                                                                                                      |
| R3-2b | 4      |            | https://suno.com/s/PcpgQ6uSweSMhsgZ | Style Influence 90%. Тут полностью halftime/brostep получился. Тут оригинальные вставки есть. Т.е оригинальный риф узнаётся.                                                                                                                                                    |
| J1-1a | 4      |            | https://suno.com/s/fsNuVt0br3705Y5W |                                                                                                                                                                                                                                                                                 |
| J1-1b | 5      |            | https://suno.com/s/meAlg2QcouklEraS | Более выигрышно звучит у остальных J1 есть классные моменты но не во всём треке. Лид понравился                                                                                                                                                                                 |
| J1-2a | 4      |            | https://suno.com/s/IpK0rZW21tM4Sqd7 |                                                                                                                                                                                                                                                                                 |
| J1-2b | 4      |            | https://suno.com/s/zFtEP1u62tISjPci |                                                                                                                                                                                                                                                                                 |
| J2-1a | 4      |            | https://suno.com/s/jvvEpM5Vlcuudof3 | У всех j2 вроде есть задорное настроение, сильно выделяются brass вроде есть схожесть с sega с sonic отдалённо, но brass всё превращат не знаю как описать во что. Кажется brass нужно заменять на то что в sega процессоре заменяли                                            |
| J2-1b | 4      |            | https://suno.com/s/27XaIcMqulLLUZ0A |                                                                                                                                                                                                                                                                                 |
| J2-2a | 4      |            | https://suno.com/s/kIjWoxUV0WThm6sh |                                                                                                                                                                                                                                                                                 |
| J2-2b | 4      |            | https://suno.com/s/GNjmbBXw8QcYHI2B |                                                                                                                                                                                                                                                                                 |
| J3-1a | 4      |            | https://suno.com/s/6LJjgsIWG7XkQ0NY | Всё классно, но тут живые струнные, мне не нравится. А настроение минорное, хорошее. Наверно это струнный                                                                                                                                                                       |
| J3-1b | 4      |            | https://suno.com/s/I7Tza6MFwrnQBFVz | Из за струнных 4, а так хороший трек                                                                                                                                                                                                                                            |
| J3-2a | 4      |            | https://suno.com/s/UAeR6Z0TOgv1yXfw | Из за струнных 4, а так хороший трек                                                                                                                                                                                                                                            |
| J3-2b | 4      |            | https://suno.com/s/fLs4xG3WoMfoHXah | Из за струнных 4, а так хороший трек                                                                                                                                                                                                                                            |
| J4-1a | 4      |            | https://suno.com/s/84DWtmDcF4BRMUbk | Немного живых труб опять влезло, а так трек почти на 5 тянет                                                                                                                                                                                                                    |
| J4-1b | 4      |            | https://suno.com/s/IxQ5fppqAiLsoh19 | Эта почему-то напряжённая, тут как-будто лид заменяет вокал, не понял хорошо это или нет. Трубы не нравятся.                                                                                                                                                                    |
| J4-2a | 4      |            | https://suno.com/s/geCU6fsvPnkVsgiD | Опять трубы, из видимо исключать надо.                                                                                                                                                                                                                                          |
| J4-2b | 4      |            | https://suno.com/s/5W4iv2WCSDaB0pGh | Трубы-нет, лид классный, можно развивать. Но трек меньше поход на DnB кроме Breakdown                                                                                                                                                                                           |
| J5-1a | 4      |            | https://suno.com/s/WaDvtlJX3CJuvnas | В части neurofunk всё терпимо, разнообразит нормально. Почему-то похоже на J1 ходя это не кавер                                                                                                                                                                                 |
| J5-1b | 4      |            | https://suno.com/s/zR2VgtVlDCGvVNA9 | Тут больше нейрофанка                                                                                                                                                                                                                                                           |
| J5-2a | 4      |            | https://suno.com/s/87P38naqV5FE8d1j | Неплохой трек, но не понятны ожидания.                                                                                                                                                                                                                                          |
| J5-2b | 4      |            | https://suno.com/s/RisAe5uAOhRr0Xtt | Тоже похоже получилось на J1 но нейрофанк , и это не ремикс. Трек бы взял в альбом, но я не понял что я от него хочу                                                                                                                                                            |
| J6-1a | 4      |            | https://suno.com/s/HOzWGPhR9ZSYOYMz | Идея интересная, но куча труб опять                                                                                                                                                                                                                                             |
| J6-1b | 4      |            | https://suno.com/s/SVVoVr8XtIHIL7EX | Идея интересная, но куча труб опять                                                                                                                                                                                                                                             |
| J6-2a | 4      |            | https://suno.com/s/YZSOphv8hB5Fnaeg | Идея интересная, но куча труб опять                                                                                                                                                                                                                                             |
| J6-2b | 4      |            | https://suno.com/s/DxP0LsIJYBgXVjtN | дея интересная, но куча труб опять                                                                                                                                                                                                                                              |

### Примечание
Треки получились классные, я смело из могу хоть все брать на мероприятие, но для альбома ещё по выбираю. Для каждого
трека нужно придумать ожидание, чтобы слушать и понимать близко мы к тому что нужно или нет. Во всех треках на, самом
деле трубы меня смущают. То что fanfare, brass. Это вопрос вкуса, я не любитель. Смущают струнные. Такие мысли меня
посетили, каждый J трек это может быть какая-то знаковая игра на знаковой консоли.
- J1 - NES - подумать
- J2 - Sega - Хотелось бы sonic
- J3 - боевой настрой хорошо, скрипки, плохо, нужно выбрать референс консоль и игра
- J4 - похоже на музыку из игры, ближе к sonic, но тут всё равно влезли живые трубы. Попробовать сделать по настроению
  jrpg. В j4 мелодики много
- J5 - надо придумать - примешать нейрофанк кажется ок.
- J6 - надо придумать

### Залп 2 — ядро, второй заход

Колонка «Попал в ожидание» — да / частично / нет, по строке «Ожидание» под промптом. Оценка 1–5 — как раньше.

| Код   | Оценка | Попал в ожидание | Что не так | Ссылка | Примечания |
|-------|--------|------------------|------------|--------|------------|
| J1-1a |        |                  |            |        |            |
| J1-1b |        |                  |            |        |            |
| J1-2a |        |                  |            |        |            |
| J1-2b |        |                  |            |        |            |
| J1-3a |        |                  |            |        |            |
| J1-3b |        |                  |            |        |            |
| J2-1a |        |                  |            |        |            |
| J2-1b |        |                  |            |        |            |
| J2-2a |        |                  |            |        |            |
| J2-2b |        |                  |            |        |            |
| J2-3a |        |                  |            |        |            |
| J2-3b |        |                  |            |        |            |
| J2-4a |        |                  |            |        |            |
| J2-4b |        |                  |            |        |            |
| J3-1a |        |                  |            |        |            |
| J3-1b |        |                  |            |        |            |
| J3-2a |        |                  |            |        |            |
| J3-2b |        |                  |            |        |            |
| J3-3a |        |                  |            |        |            |
| J3-3b |        |                  |            |        |            |
| J4-1a |        |                  |            |        |            |
| J4-1b |        |                  |            |        |            |
| J4-2a |        |                  |            |        |            |
| J4-2b |        |                  |            |        |            |
| J4-3a |        |                  |            |        |            |
| J4-3b |        |                  |            |        |            |
| J4-4a |        |                  |            |        |            |
| J4-4b |        |                  |            |        |            |
| J5-1a |        |                  |            |        |            |
| J5-1b |        |                  |            |        |            |
| J5-2a |        |                  |            |        |            |
| J5-2b |        |                  |            |        |            |
| J5-3a |        |                  |            |        |            |
| J5-3b |        |                  |            |        |            |
| J6-1a |        |                  |            |        |            |
| J6-1b |        |                  |            |        |            |
| J6-2a |        |                  |            |        |            |
| J6-2b |        |                  |            |        |            |
| J6-3a |        |                  |            |        |            |
| J6-3b |        |                  |            |        |            |

### Залп 2 — соседи

Колонка «Попал в ожидание» — по строке под промптом. N4 закрыт репризой R3, строк нет.

| Код   | Оценка | Попал в ожидание | Что не так | Ссылка | Примечания |
|-------|--------|------------------|------------|--------|------------|
| N1-1a |        |                  |            |        |            |
| N1-1b |        |                  |            |        |            |
| N1-2a |        |                  |            |        |            |
| N1-2b |        |                  |            |        |            |
| N2-1a |        |                  |            |        |            |
| N2-1b |        |                  |            |        |            |
| N2-2a |        |                  |            |        |            |
| N2-2b |        |                  |            |        |            |
| N3-1a |        |                  |            |        |            |
| N3-1b |        |                  |            |        |            |
| N3-2a |        |                  |            |        |            |
| N3-2b |        |                  |            |        |            |
| N5-1a |        |                  |            |        |            |
| N5-1b |        |                  |            |        |            |
| N5-2a |        |                  |            |        |            |
| N5-2b |        |                  |            |        |            |
| B1-1a |        |                  |            |        |            |
| B1-1b |        |                  |            |        |            |
