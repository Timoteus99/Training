# think of a number between 1 and a 1000, then run the code.
def get_one_of(prompt, values):
    """
    Prompt for a response in values,
    return response string
    """
    while True:
        value = input(prompt).strip().lower()
        if value in values:
            return value

lo, hi = 1, 1000
got_it = False
for attempt in range(10):
    guess = (lo + hi) // 2
    print("I guess {}!".format(guess))
    res = get_one_of("Was this [L]ow, [H]igh, or [C]orrect? ", {"l", "h", "c"})            
    if res == "l":
        lo = guess + 1
    elif res == "h":
        hi = guess - 1
    else:  # correct!
        got_it = True
        break
    if lo > hi:
        break
if got_it:
    print("Got it in {} guesses!".format(attempt + 1))
else:
    print("Ne me zajebavat...")