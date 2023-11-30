import time
from prime_number import calc_primes


start_time = time.time()  # 処理時間計算用

ranges = [  # 素数を計算する開始値と終了値
    (1_000_000, 1_600_000),
    (2_000_000, 2_900_000),
    (3_000_000, 3_300_000),
]
for num_start, num_end in ranges:
    calc_primes(num_start, num_end)

end_time = time.time()  # 終了時間
print(f"Processing time: {end_time - start_time} seconds")

