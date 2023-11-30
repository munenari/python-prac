import time
import threading

import _xxsubinterpreters as interpreters

start_time = time.time()  # 処理時間計算用

## interpreters.run_string()に渡す関数を文字列として読み込む
with open("prime_number.py", "r") as f:
    # ファイルの内容を読み込む
    calc_prime_script = f.read()

# 数値範囲のリスト
num_ranges = [
    "1_000_000, 1_600_000",
    "2_000_000, 2_900_000",
    "3_000_000, 3_300_000",
]

# サブインタープリターのIDとスレッドを格納するリスト
intp_ids = []
threads = []

# 各サブインタープリターの作成とスレッドの開始
for num_range in num_ranges:
    intp_id = interpreters.create()
    intp_ids.append(intp_id)

    thread = threading.Thread(
        target=interpreters.run_string,
        args=(intp_id, calc_prime_script + f"\ncalc_primes({num_range})"),
    )
    thread.start()
    threads.append(thread)

# すべてのスレッドが完了するのを待つ
for thread in threads:
    thread.join()

# すべてのサブインタープリターを破棄
for intp_id in intp_ids:
    interpreters.destroy(intp_id)

end_time = time.time()
print(f"Processing time: {end_time - start_time} seconds")
