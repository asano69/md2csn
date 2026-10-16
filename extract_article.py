import bz2
import re
import xml.etree.ElementTree as ET

INDEX_FILE = "jawiki-latest-pages-articles-multistream-index.txt"
DUMP_FILE = "jawiki-latest-pages-articles-multistream.xml.bz2"

TARGET_TITLE = "日本語"
OUTPUT_FILE = "japanese.txt"


def find_index(index_file, target_title):
    """タイトルに一致するページのオフセットとページIDを探す。"""
    with open(index_file, "r", encoding="utf-8") as f:
        for line in f:
            parts = line.rstrip("\n").split(":", 2)
            if len(parts) == 3 and parts[2] == target_title:
                return int(parts[0]), parts[1]

    raise ValueError(f"記事が見つかりません: {target_title}")


def read_bz2_stream(dump_file, offset):
    """指定オフセットから、1つの bzip2 ストリームを展開する。"""
    decompressor = bz2.BZ2Decompressor()
    chunks = []

    with open(dump_file, "rb") as f:
        f.seek(offset)

        while not decompressor.eof:
            compressed = f.read(1024 * 1024)
            if not compressed:
                raise EOFError("圧縮ストリームが途中で終了しました")

            chunks.append(decompressor.decompress(compressed))

    return b"".join(chunks)


def extract_page(xml_bytes, target_title, target_id):
    """ストリーム内から対象ページの本文を取り出す。"""
    xml_text = xml_bytes.decode("utf-8")

    # 1つのストリームに含まれる各ページを探す
    for match in re.finditer(
        r"<page(?:\s[^>]*)?>.*?</page>",
        xml_text,
        re.DOTALL,
    ):
        page_xml = match.group(0)

        try:
            page = ET.fromstring(page_xml)
        except ET.ParseError:
            continue

        title = page.findtext("title")
        page_id = page.findtext("id")

        if title != target_title or page_id != str(target_id):
            continue

        # 最新リビジョンの本文を取得
        revisions = page.findall("revision")
        if not revisions:
            raise ValueError("リビジョンが見つかりません")

        text_element = revisions[-1].find("text")
        if text_element is None:
            raise ValueError("本文が見つかりません")

        return text_element.text or ""

    raise ValueError(
        f"ストリーム内にページが見つかりません: {target_title} (ID={target_id})"
    )


def main():
    offset, page_id = find_index(INDEX_FILE, TARGET_TITLE)
    print(f"タイトル: {TARGET_TITLE}")
    print(f"オフセット: {offset}")
    print(f"ページID: {page_id}")

    xml_bytes = read_bz2_stream(DUMP_FILE, offset)
    content = extract_page(xml_bytes, TARGET_TITLE, page_id)

    with open(OUTPUT_FILE, "w", encoding="utf-8") as f:
        f.write(content)

    print(f"本文を保存しました: {OUTPUT_FILE}")
    print(f"文字数: {len(content):,}")
    print("\n--- 本文の先頭500文字 ---")
    print(content[:500])


if __name__ == "__main__":
    main()
