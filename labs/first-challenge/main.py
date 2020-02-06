word = input()

substrings = set()
new_word = ""
longest = ""
start = 0
position = 0

while position < len(word):

  old_length = len(substrings)

  substrings.add(word[position])

  new_length = len(substrings)

  if old_length == new_length or position + 1 == len(word):

    print(new_word)

    if len(new_word) > len(longest):

      longest = new_word

    start = start + 1

    position = start

    new_word = ""

    substrings = set()

  else:

    new_word += word[position]

    position += 1



print (len(longest))
