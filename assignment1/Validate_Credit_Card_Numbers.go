# Enter your code here. Read input from STDIN. Print output to STDOUT

import re

def validate_cards(card_numbers):
    
    pattern = r'^[456]\d{3}-?\d{4}-?\d{4}-?\d{4}$'
    repetition_pattern = r'(.)\1{3}'
    
    for card_number in card_numbers:
        card_number2 = re.sub('-', '', card_number)
        if re.match(pattern, card_number) and not re.search(repetition_pattern, card_number2):
            print("Valid")
        else:
            print("Invalid")
    
    return None

N = int(input().strip())
card_numbers = []
for index in range(N):
    card_numbers.append(input().strip())

validate_cards(card_numbers)

