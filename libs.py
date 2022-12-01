import math


def format_runtime(ms):
    # Microseconds
    if ms <= 1:
        return f"{round(ms * 1000)}µs"
    # Milliseconds
    if ms < 1000:
        whole_ms = math.floor(ms)
        rem_ms = ms - whole_ms
        return f"{whole_ms}ms " + format_runtime(rem_ms)
    sec = ms / 1000
    # Seconds
    if sec < 60:
        whole_sec = math.floor(sec)
        rem_ms = ms - whole_sec * 1000
        return f"{whole_sec}s " + format_runtime(rem_ms)
    # Minutes (hopefully it doesn't get to this point lol)
    return f"{math.floor(sec / 60)}m " + format_runtime((sec % 60) * 1000)
