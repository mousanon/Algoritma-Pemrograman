nilai = int(input("Nilai: "))
budget = 0

if nilai > 80:
    print("OP")
    if nilai == 100:
        print("II")
        budget += 100000
        print(budget)
    
print(nilai)
print(budget)