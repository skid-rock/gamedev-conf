# DnB для конференции — поджанры и промпты для Suno v6

Продолжение раздела «Party — DnB × чиптюн» из файла 03. Там один усреднённый DnB. Здесь — поджанры по отдельности, чтобы понять, какой из них «наш», и под какую площадку какой подходит.

Не тестировалось. Правила промптов — из файла 02. Общие настройки те же, что в 03: Custom Mode, Instrumental, Style Influence 80%, Weirdness 30%, Variety Normal на поиске, 0 — на доводке.

Общее про DnB в Suno: жанр держится на барабанах, поэтому в Style барабаны описываем первыми после жанра и настроения — тип брейка, характер снейра, плотность. BPM всегда 170–176, иначе модель уплывает в брейкбит или хаус.

## Карта поджанров

| Поджанр | Энергия | Настроение | Где на конференции |
|---|---|---|---|
| Liquid | средняя | тёплое, мелодичное | вечерний приём, лаунж после 18:00, промо-ролик «про людей» |
| Jungle | высокая | сырое, олдскульное | афтепати, гейм-джем |
| Atmospheric / intelligent | низкая-средняя | космос, глубина | деловая зона к вечеру, фон для демо-стендов |
| Dancefloor / rollers | высокая | ровный драйв | обратный отсчёт, разогрев перед открытием |
| Jump-up | очень высокая | хулиганское | афтепати, гейм-джем на финише |
| Neurofunk | высокая | техно-агрессия | тизер, трейлер, «боссфайт»-момент церемонии |
| Techstep / darkstep | высокая | мрачное, индустриальное | промо для хоррор/шутер-секции, если такая есть |
| Halftime | средняя | тяжёлое, замедленное | перебивка «дорого и модно», партнёрские ролики |
| Drumfunk | средняя | нервное, джазовое | фирменная фишка для тех, кто разбирается |
| Minimal / deep | низкая | гипноз | фон для позднего нетворкинга |
| Breakcore / drill'n'bass | очень высокая | хаос | только как шутка в ролике, 10–15 секунд |
| Sambass / latin DnB | высокая | праздник | закрытие, финальная вечеринка |

Рекомендация: начать с трёх — **liquid**, **jungle**, **dancefloor**. Они закрывают три разных сценария и хорошо отличаются друг от друга на слух. Остальное — по результатам.

---

## 1. Liquid (liquid funk)

Мелодичный DnB: тёплые пэды, фортепиано или родес, вокальные обрывки как текстура, мягкий бас. Барабаны ровные, без агрессии.

Референсы: Calibre, High Contrast (ранний), London Elektricity, Netsky (первые два альбома), Logistics, Hospital Records в целом.

**Модель:** v6. Weirdness 25%. Duration Custom 4:00.

**Style:**
```
liquid drum and bass, warm and uplifting, smooth rolling breakbeat, soft crisp snare, deep round sub bass, rhodes chords, lush string pads, soft piano melody, chopped vocal texture used as an instrument, clean polished Hospital Records style mix, 174 BPM
```

**Lyrics:**
```
[Intro: rhodes chords and string pad, no drums, main motif on piano]
[Build: hi-hats and soft snare roll enter]
[Drop: full rolling breakbeat and sub bass, piano motif on top]
[Breakdown: drums out, pads and chopped vocal texture]
[Drop: full beat returns, rhodes and piano together]
[Outro: drums fade, piano plays the motif alone, soft end]
```

**Exclude Styles:** `lyrics, sung words, rap, aggressive bass, distortion`

Заметка: «chopped vocal texture» может вытянуть полноценный вокал. Если так — убрать и оставить только инструменты.

---

## 2. Jungle (oldschool)

Корни жанра, 1993–1996: рваные amen-брейки, регги-бас, сэмплы сирен, рэгга-выкрики, рейв-стабы. Звучит сыро и самодельно — это часть стиля.

Референсы: Goldie «Timeless» (ранние вещи), LTJ Bukem (лёгкая сторона), Shy FX «Original Nuttah», Dillinja, Remarc, Congo Natty, Aphrodite.

**Модель:** v6, второй прогон v6-wild. Weirdness 40%. Duration Custom 3:30.

**Style:**
```
oldschool jungle, raw and hyped, chopped amen break, fast rolling drums, deep reggae sub bass, rave stabs, air horn and siren samples, timestretched vocal shouts, dusty sampler sound, 1994 pirate radio production, 170 BPM
```

**Lyrics:**
```
[Intro: rave stab and siren, then amen break cuts in immediately]
[Drop: chopped amen break and reggae sub bass, rave stabs]
[Breakdown: drums drop out, timestretched pad and air horn]
[Drop: amen break returns with double-time chops, bass heavier]
[Outro: drums stop dead, one siren, end]
```

**Exclude Styles:** `lyrics, sung verses, rap verses, clean modern mix, trap`

Для чиптюн-гибрида: добавить в Style `8-bit lead melody` после стабов. Jungle и чиптюн из одной эпохи сэмплеров, сочетание естественное.

---

## 3. Atmospheric / intelligent DnB

Глубокие пэды, длинные интро, барабаны мягче, чем в jungle, бас округлый. Музыка для слушания, не для танца.

Референсы: LTJ Bukem «Logical Progression», Good Looking Records, Photek (спокойные вещи), Omni Trio, Seba, Blu Mar Ten.

**Модель:** v6. Max Mode вкл. Weirdness 25%. Duration Custom 5:00.

**Style:**
```
atmospheric drum and bass, deep and spacious, smooth rolling breakbeat, soft snare, warm round sub bass, wide ambient pads, distant piano, gentle synth arpeggio, long reverb tails, clean 90s Good Looking Records style, 172 BPM
```

**Lyrics:**
```
[Intro: ambient pads and distant piano, long, no drums]
[Build: hi-hats enter, then soft breakbeat]
[Main: rolling breakbeat and sub bass under pads, arpeggio plays the motif]
[Break: drums out, pads only, piano motif slower]
[Main: breakbeat returns, long steady groove]
[Outro: drums fade first, pads last, slow fade]
```

**Exclude Styles:** `vocals, singing, rap, aggressive bass, distortion, build-up to big drop`

---

## 4. Dancefloor / rollers

Ровный, накатывающий DnB без резких поворотов. Один грув на весь трек, бас качает, мелодии минимум. То, что играют в клубах в прайм-тайм.

Референсы: Break, Serum, Bladerunner, DLR, Alix Perez (ранний), Total Science, S.P.Y.

**Модель:** v6. Weirdness 25%. Duration Custom 4:00.

**Style:**
```
dancefloor drum and bass roller, driving and steady, tight punchy breakbeat, sharp snare, deep rolling sub bass, minimal stabs, short vocal chop, clean club mix, relentless groove, 174 BPM
```

**Lyrics:**
```
[Intro: filtered breakbeat, bass hinted, 16 bars]
[Drop: full breakbeat and rolling sub bass, minimal stab]
[Roll: same groove continues, subtle variation in hats]
[Breakdown: bass out, drums keep rolling, short vocal chop]
[Drop: bass returns, full groove to the end]
[Outro: drums cut to a single hit, end]
```

**Exclude Styles:** `lyrics, sung words, rap, melody-heavy, orchestral`

Применение: обратный отсчёт «до начала 5 минут». Ровный грув без пиков — можно резать в любом месте.

---

## 5. Jump-up

Простые, наглые, «квакающие» басы, хуковый рифф, барабаны прямые. Жанр про веселье, не про тонкости.

Референсы: Hazard, Macky Gee, Sub Focus (ранний), DJ Hype, Turno, Bou.

**Модель:** v6-wild. Weirdness 45%. Duration Custom 3:00.

**Style:**
```
jump up drum and bass, cheeky and hyped, punchy two-step breakbeat, snappy snare, wobbling talking bass riff, big bass hook, arcade sound effects, crowd hype energy, loud club mix, 175 BPM
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

Гибрид с чиптюном напрашивается: «talking bass» и 8-bit-лид — оба хулиганские. Добавить `8-bit square lead doubles the bass riff`.

---

## 6. Neurofunk

Технологичный, холодный DnB: бас с постоянно меняющимся тембром, ритм с синкопами, звук как из научной фантастики. Самый «продюсерский» поджанр.

Референсы: Noisia, Phace, Misanthrop, Mefjus, Black Sun Empire, Ed Rush & Optical «Wormhole».

**Модель:** v6. Style Influence 85%. Weirdness 30%. Duration Custom 3:30.

**Style:**
```
neurofunk drum and bass, dark and precise, tight syncopated breakbeat, metallic snare, growling modulated reese bass with evolving texture, sci-fi synth stabs, cold industrial atmosphere, surgical modern mix, 174 BPM
```

**Lyrics:**
```
[Intro: cold sci-fi pad and single sub pulse, no drums]
[Build: mechanical percussion and rising noise sweep]
[Drop: syncopated breakbeat and modulated reese bass, stabs answer the bass]
[Breakdown: drums out, bass texture morphs alone]
[Drop: full beat, bass pattern more complex]
[Outro: everything cuts to a low drone, end]
```

**Exclude Styles:** `vocals, singing, rap, warm pads, piano, happy`

Применение: тизер конференции в стиле трейлера, момент «объявление победителя» с драматизацией. В фон не годится, слишком требует внимания.

---

## 7. Techstep / darkstep

Предшественник neurofunk, конец 90-х: мрачно, индустриально, барабаны жёстче и проще, бас грязный. Меньше полировки, больше давления.

Референсы: Ed Rush & Optical, Bad Company UK, Konflict, Technical Itch, Dom & Roland, Dylan.

**Модель:** v6. Weirdness 35%. Duration Custom 3:30.

**Style:**
```
techstep drum and bass, dark and industrial, hard two-step breakbeat, heavy snare, distorted reese bass, metallic percussion hits, ominous sci-fi pads, gritty late 90s production, 172 BPM
```

**Lyrics:**
```
[Intro: ominous pad and metallic hits, no drums]
[Drop: hard two-step breakbeat and distorted reese, sparse stabs]
[Break: drums out, bass drone and pad]
[Drop: beat returns, extra percussion layer]
[Outro: drums stop, pad and one final hit]
```

**Exclude Styles:** `vocals, singing, rap, bright melody, piano, uplifting`

---

## 8. Halftime

Барабаны звучат вдвое медленнее (снейр на 3-ю долю, а не на 2-ю и 4-ю), при том же темпе 170+. Ощущение тяжёлого, замедленного грува — родня трэпа и дабстепа.

Референсы: Ivy Lab, Alix Perez (поздний), Noisia (halftime-вещи), Halogenix, Monty, лейбл 1985 Music.

**Модель:** v6. Weirdness 30%. Duration Custom 3:00.

**Style:**
```
halftime drum and bass, heavy and hypnotic, slow half-time drum pattern at fast tempo, hard snare on the three, deep sub bass, wonky bass stabs, sparse dark synth, spacious mix with lots of air, 172 BPM
```

**Lyrics:**
```
[Intro: sub bass and single hi-hat pattern, no snare]
[Drop: half-time drums, deep sub and wonky bass stabs]
[Variation: extra percussion, bass pattern changes]
[Break: drums out, sub drone]
[Drop: half-time beat returns, heavier]
[Outro: drums stop, sub fades]
```

**Exclude Styles:** `vocals, singing, rap, fast rolling drums, full-time breakbeat`

Риск: модель может прочитать «half-time» как темп 86 BPM и выдать трэп. Если так — вторая попытка со Style Influence 90% и фразой `drum and bass tempo` в начале.

Применение: партнёрские ролики, где нужно «дорого и современно» без суеты.

---

## 9. Drumfunk / choppage

Барабаны — главный инструмент. Брейки нарезаны так мелко, что ритм постоянно меняется, при этом грув держится. Много джазовых сэмплов.

Референсы: Paradox, Fanu, Fracture & Neptune, Breakage (ранний), Seba, Equinox.

**Модель:** v6-wild. Weirdness 50%. Duration Custom 4:00.

**Style:**
```
drumfunk, intricate and jazzy, heavily chopped and rearranged live breakbeats, constantly shifting drum patterns, ghost notes, deep warm sub bass, jazz double bass sample, rhodes stabs, muted trumpet phrase, dusty vinyl texture, 170 BPM
```

**Lyrics:**
```
[Intro: solo chopped breakbeat, no bass, drums keep rearranging]
[Main: sub bass and rhodes stabs join, drums stay complex]
[Variation: muted trumpet phrase, drums thin out to hi-hats and ghost notes]
[Main: full chopped breaks return, double bass sample]
[Outro: drums alone again, then stop]
```

**Exclude Styles:** `vocals, singing, rap, straight simple drum loop, EDM`

Самый рискованный для генерации: ИИ-модели тяготеют к ровному лупу. Считать удачей, если хотя бы одна из четырёх генераций даёт подвижные барабаны.

---

## 10. Minimal / deep (autonomic)

Разреженный DnB: мало элементов, много воздуха, барабаны тихие и сухие, бас глубокий. Настроение — ночь, поздний час.

Референсы: dBridge & Instra:mental (серия Autonomic), ASC, Consequence, Synkro, лейбл Exit Records.

**Модель:** v6. Weirdness 25%. Duration Custom 5:00.

**Style:**
```
minimal deep drum and bass, hypnotic and late night, sparse dry drums, quiet snare, deep sub bass, subtle detuned synth pad, lots of empty space, dubby delays, autonomic style, 170 BPM
```

**Lyrics:**
```
[Intro: detuned pad and dub delay, no drums, long]
[Main: sparse dry drums and deep sub enter quietly]
[Section: same groove, pad slowly shifts, no climax]
[Break: drums out, delay tails only]
[Main: drums return, still quiet]
[Outro: everything fades slowly]
```

**Exclude Styles:** `vocals, singing, rap, big drop, loud, stabs, climax`

Применение: поздний нетворкинг, когда людям надо разговаривать, но эмбиент уже слишком сонный.

---

## 11. Breakcore / drill'n'bass

Не DnB в строгом смысле, но родня: те же брейки, порезанные до хаоса, темп плавает, мелодии издевательские. Как акцент — работает, как фон — нет.

Референсы: Squarepusher (Hard Normal Daddy, Big Loada), Aphex Twin (Come to Daddy, drill-вещи), Venetian Snares, Bogdan Raczynski, Igorrr.

**Модель:** v6-wild. Weirdness 65%. Duration Custom 0:20.

**Style:**
```
drill and bass breakcore burst, chaotic and playful, hyper-chopped amen break, stuttering glitch edits, rubbery acid bass, 8-bit melody fragment, no intro build-up, single short burst, abrupt stop, 180 BPM
```

**Lyrics:**
```
[Burst: chopped amen and acid bass explode immediately, 8-bit melody fragment]
[End: abrupt stop on a glitch]
```

**Exclude Styles:** `vocals, singing, rap, slow intro, ambient, steady groove`

Применение: 10–15 секунд в промо-ролике на монтажной склейке. В остальном — нет.

---

## 12. Sambass / latin DnB

Бразильская ветка: перкуссия самбы, духовые, гитары босса-новы поверх DnB-барабанов. Праздничный звук.

Референсы: DJ Marky, Patife, XRS, Drumagick, Fernanda Porto «Sambassim», Bungle.

**Модель:** v6. Weirdness 35%. Duration Custom 4:00.

**Style:**
```
brazilian sambass drum and bass, festive and sunny, rolling breakbeat, samba percussion layer, cuica and shakers, bossa nova nylon guitar, bright brass section, deep warm sub bass, live band feel, 174 BPM
```

**Lyrics:**
```
[Intro: nylon guitar and shakers, no drums]
[Build: samba percussion enters, then breakbeat]
[Drop: full breakbeat, sub bass and brass section carry the motif]
[Breakdown: drums out, guitar and cuica]
[Drop: everything back, percussion busier]
[Outro: brass hit, percussion tail, end]
```

**Exclude Styles:** `lyrics, sung words, rap, dark, industrial`

Применение: закрытие конференции, финальная вечеринка, ролик с итогами.

---

## Гибриды с чиптюном — как встроить в любой поджанр

Промпт из файла 03 («Party») — это, по сути, dancefloor DnB + 8-bit. Чтобы получить чиптюн-версию любого поджанра выше, добавлять в Style после басового описания один из блоков:

- Лёгкий акцент: `8-bit arpeggio accents`
- Лид: `8-bit square wave lead carries the main melody`
- Перкуссия: `arcade sound effects used as percussion, coin and jump sounds`
- Полный: `chiptune lead, NES-style arpeggios, arcade sound effects as percussion`

Лучше всего чиптюн ложится на jungle (эпоха), jump-up (хулиганство) и dancefloor (простота). Хуже всего — на liquid и minimal, где он спорит с настроением.

## Цепочка Cover

Как в файле 03: выбрать главную тему → Cover с любым Style из этого файла → проверить, узнаётся ли мотив. Для DnB мотив лучше всего выживает в liquid и atmospheric (есть место для мелодии), хуже всего — в jump-up и neurofunk (бас забирает всё внимание).

## Журнал

| Дата | Поджанр | Модель | Что менял в промпте | Ссылка | Оценка 1–5 | Что не так |
|---|---|---|---|---|---|---|
| | | | | | | |
