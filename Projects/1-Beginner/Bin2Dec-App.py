binary = input("Enter binary number: ") # taking user input
index = 0 # variables to count index of current digit and total decimal number
total = 0
try:  # try-except block to ensure input is binary 
	for i in reversed(binary): # input string is reversed so we increase power of 2 from right to left
		current = int(i) # converting input string to int, raises exception in case of non-digit character
		if current != 0 and current != 1: 
			raise ValueError() # raising exception in case of non-binary input
		if current == 1: # incrementing total by 2^index at every occurence of a 1 in the input
			total += 2**index
		index += 1 # incrementing index as we parse
	print(total) # printing total
except: # catching exceptions raised and printing message
	print("Invalid input: only 0's and 1's are accepted ")