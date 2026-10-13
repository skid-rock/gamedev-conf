#!/bin/bash
# Экспорт презентаций Keynote (.key) в .pptx силами самого Keynote.
#
# Формат .key закрытый, разобрать его сторонними средствами нельзя,
# поэтому конвертацию делает само приложение через AppleScript.
#
# Использование:
#   ./key2pptx.sh КАТАЛОГ_НАЗНАЧЕНИЯ файл.key [файл.key ...]
#
# Keynote адресуется по идентификатору com.apple.iWork.Keynote, а не по имени:
# имя "Keynote" может перехватить постороннее приложение с таким же
# исполняемым файлом, и обращение уйдёт не туда.

set -u

if [ $# -lt 2 ]; then
	echo "использование: $0 КАТАЛОГ_НАЗНАЧЕНИЯ файл.key [файл.key ...]" >&2
	exit 2
fi

outdir="$1"; shift
mkdir -p "$outdir" || exit 1

ok=0
fail=0
for src in "$@"; do
	if [ ! -e "$src" ]; then
		echo "  ! нет файла: $src" >&2
		fail=$((fail + 1))
		continue
	fi
	base="$(basename "$src")"
	name="${base%.key}"
	dst="$outdir/$name.pptx"
	abs_src="$(cd "$(dirname "$src")" && pwd)/$base"
	abs_dst="$(cd "$outdir" && pwd)/$name.pptx"

	printf 'экспортирую %s ... ' "$name"
	err=$(osascript - "$abs_src" "$abs_dst" <<'APPLESCRIPT' 2>&1
on run argv
	set inFile to POSIX file (item 1 of argv)
	set outFile to POSIX file (item 2 of argv)
	tell application id "com.apple.iWork.Keynote"
		set doc to open inFile
		export doc to outFile as Microsoft PowerPoint
		close doc saving no
	end tell
end run
APPLESCRIPT
	)
	if [ $? -eq 0 ] && [ -s "$abs_dst" ]; then
		printf 'готово (%s)\n' "$(du -h "$dst" | cut -f1 | tr -d ' ')"
		ok=$((ok + 1))
	else
		printf 'ОШИБКА\n'
		[ -n "$err" ] && echo "    $err" >&2
		fail=$((fail + 1))
	fi
done

echo
echo "экспортировано: $ok, с ошибкой: $fail"
[ "$fail" -gt 0 ] && exit 1
exit 0
