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
| Exclude Styles | `lyrics, sung words, rap, vocals, singing, soft pads, piano, brass, trumpets, horns, fanfare, orchestral, strings, live instruments` | то же плюс по промпту | по промпту |

Слово «instrumental» и «no ...» в Style не пишем — тумблер и Exclude. Формат Style — как в файле 04. Suno у себя
переписывает промпт в форму с двоеточием (так он выглядит в json из gamepush), это не наша
забота [допущение: в json сохранён переписанный текст, а вводилась форма файла 04].

## План на 21 час — три залпа

| Залп | Что                                                      | Пар   | Когда                |
|------|----------------------------------------------------------|-------|----------------------|
| 1    | R1, R2, R3 по одной паре + J1–J6 по две пары             | 15    | сразу                |
| 2    | N1–N5 по две пары + доработка 2–3 лучших J по одной паре | 12–14 | после оценок залпа 1 |
| 3    | добор дыр в трек-листе                                   | 4–6   | после оценок залпа 2 |

После залпа 3: отбор 10–11, порядок, названия.

Оценки — в чат, одной строкой на трек: `код — оценка 1–5 — что не так`. Понравившееся — качать в `gamepush/` через
suno-download с `--name '{artist} - {title}'`, где title — рабочий код. Журнал в конце файла веду я.

## Целевой трек-лист (план, не порядок)

| Код | Что                        | Поджанр     | Ось разнообразия                                                    |
|-----|----------------------------|-------------|---------------------------------------------------------------------|
| J1  | базовый, контрольная точка | jump-up     | NES-лид дублирует бас, наглый                                       |
| J2  | FM-версия                  | jump-up     | Mega Drive slap bass, металлический FM-лид, ночной клубный тон      |
| J3  | военный                    | jump-up     | марш-стаккато, гармонический минор, взрывы                          |
| J4  | аркадный                   | jump-up     | Neo Geo дудки, лид ведёт свою мелодию поверх баса                   |
| J5  | с приправой neuro          | jump-up     | бас с меняющимся тембром, two-step и 8-bit на месте                 |
| J6  | солнечный                  | jump-up     | мажор, бонус-уровень, фанфары «stage clear»                         |
| N1  | jungle 8-bit               | jungle      | amen-брейк, 8-bit-мелодия                                           |
| N2  | jungle с talking bass      | jungle      | amen + jump-up-бас, мост между ядром и jungle                       |
| N3  | роллер                     | dancefloor  | ровный грув, 8-bit-лид как единственная мелодия                     |
| N4 | перелом | halftime | закрыт репризой R3: рифф J1 узнаётся в R3-2a (полный темп, halftime-вставки) и R3-2b (полностью halftime). Какую брать и брать ли — решение отложено |
| N5  | передышка                  | liquid-лайт | без пэдов и пианино, лид ведёт тему, чиптюн-арпеджио вместо пэдов   |
| B1  | шутка                      | breakcore   | 20–30 секунд, бонус-позиция, если войдёт                            |

Если после залпа 1 две оси ядра сольются на слух, одну заменяю на «тёмный» jump-up с half-time-вставкой в середине
трека.

**После залпа 1:** оси J1–J6 пересобираются как «консоль + настроение», по одной консоли на трек, чтобы у каждого было ожидание. Зафиксировано: J3 — Amiga. Остальная раскладка обсуждается в отдельной сессии, промпты J переписываются после неё.

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

## Залп 2 — соседи (черновик, правлю после оценок залпа 1)

Все N: v6, Weirdness 35%, Variety Normal, Style Influence 80%, Duration 3:45. Инварианты на месте: лид, SFX, loud club
mix.

### N1. Jungle 8-bit

Файл 04 §2 плюс 8-bit-лид, который ведёт мелодию — вторая формула хука.

**Style:**
```
oldschool jungle, raw and hyped, chopped amen break, fast rolling drums, deep reggae sub bass, rave stabs, 8-bit square lead carries the main melody, air horn and siren samples, arcade sound effects, dusty sampler sound, loud club mix, 170 BPM
```

**Lyrics:**
```
[Intro: rave stab and siren, then amen break cuts in immediately]
[Drop: chopped amen break and reggae sub bass, 8-bit lead plays the main melody]
[Breakdown: drums drop out, 8-bit lead alone with air horn]
[Drop: amen break returns with double-time chops, bass heavier, lead melody on top]
[Outro: drums stop dead, one siren, end]
```

**Exclude Styles:** `lyrics, sung verses, rap verses, vocals, singing, timestretched vocal shouts, clean modern mix, trap, brass, trumpets, horns, fanfare, orchestral, strings, live instruments`

Заметка: `timestretched vocal shouts` из файла 04 убраны в Exclude — на альбоме без вокала выкрики будут торчать. Если
jungle без них звучит пусто — вернуть в Style.

### N2. Jungle с talking bass — мост к ядру

**Style:**
```
oldschool jungle, raw and cheeky, chopped amen break, fast rolling drums, wobbling talking bass riff, 8-bit square lead doubling the bass riff, rave stabs, siren samples, arcade sound effects, loud club mix, 172 BPM
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

**Exclude Styles:** `lyrics, sung words, rap, vocals, singing, soft pads, piano, clean modern mix, brass, trumpets, horns, fanfare, orchestral, strings, live instruments`

### N3. Роллер

Файл 04 §4 с 8-bit-лидом как единственной мелодией.

**Style:**
```
dancefloor drum and bass roller, driving and steady, tight punchy breakbeat, sharp snare, deep rolling sub bass, minimal stabs, 8-bit square lead plays a short repeating motif, arcade sound effects sparse, loud club mix, relentless groove, 174 BPM
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

**Exclude Styles:** `lyrics, sung words, rap, vocals, singing, melody-heavy, soft pads, piano, brass, trumpets, horns, fanfare, orchestral, strings, live instruments`

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

**Exclude Styles:** `lyrics, sung words, rap, vocals, singing, soft pads, piano, fast rolling drums, full-time breakbeat, trap, brass, trumpets, horns, fanfare, orchestral, strings, live instruments`

### N5. Liquid-лайт — передышка

Liquid без пэдов и пианино: их место занимают чиптюн-арпеджио, лид ведёт тему.

**Style:**
```
liquid drum and bass, warm and bright, smooth rolling breakbeat, crisp snare, deep round sub bass, 8-bit square lead carries a melodic theme, plucked chiptune arpeggios instead of pads, sparse arcade sound effects, clean polished mix, 174 BPM
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

**Exclude Styles:** `lyrics, sung words, rap, vocals, singing, piano, string pads, aggressive bass, distortion, brass, trumpets, horns, fanfare, orchestral, strings, live instruments`

Заметка: `clean polished mix` вместо `loud club mix` — единственное отступление от инварианта, иначе передышки не будет.
Если трек выпадает из альбома по громкости — вернуть `loud club mix`.

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

**Exclude Styles:** `vocals, singing, rap, lyrics, slow intro, ambient, steady groove, brass, trumpets, horns, fanfare, orchestral, strings, live instruments`

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
