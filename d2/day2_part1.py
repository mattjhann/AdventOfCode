def main():
    input = []

    with open('day2_input.csv', 'r') as file:
        for line in file:
            input.append(line.replace('\n', '').split(' '))

    count = 0
    for report in input:
        if validateReport(report, True): count += 1
    
    print(count)


def validateReport(report, canRemove):
    for i in range(len(report)):
        report[i] = int(report[i])
    ascending = isAscending(report[0], report[1])

    i = 0
    while i < len(report) - 1:
        if validateNums(report[i], report[i+1], ascending) == False: 
            if canRemove:
                if i == len(report) - 2:
                    return True
                elif validateNums(report[i-1], report[i+1], ascending):
                    report.pop(i)
                elif len(report) > i + 1:
                    return False
                elif validateNums(report[i], report[i+2], ascending):
                    report.pop(i+1)
                else:
                    return False
                
                canRemove = False
                i = i-1
                continue
            else:
                return False
        
        i = i + 1 # iterate manually

    return True


def validateNums(val1, val2, asc):
    if asc != isAscending(val1, val2):
        print(f"{val1} & {val2} : should be ascending={asc}")
        return False
    if abs(val1 - val2) < 1 or abs(val1 - val2) > 3:
        print(f"{val1} & {val2} : Difference is out of bounds (<1 or >3)")
        return False
    return True


def isAscending(val1, val2):
    return True if (val1 - val2 < 0) else False


main()