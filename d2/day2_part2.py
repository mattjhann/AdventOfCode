def main():
    input = []

    with open('day2_input.csv', 'r') as file:
        for line in file:
            input.append(line.replace('\n', '').split(' '))

    count = 0
    for report in input:
        for i in range(len(report)):
            report[i] = int(report[i])

        if validateReport(report, True) or validateReport(report, False):
            count += 1
            continue
        
        res = []
        for i in range(len(report)):
            j = report.copy()
            j.pop(i)
            if (validateReport(j, True) or validateReport(j, False)):
                count += 1
                break

    print(count)


def validateReport(report, asc):
    for i in range(len(report) - 1):
        if validateNums(report[i], report[i+1], asc) == False:
            return False
    
    return True
    

def validateNums(val1, val2, asc):
    if asc != isAscending(val1, val2):
        return False
    if abs(val1 - val2) < 1 or abs(val1 - val2) > 3:
        return False
    return True


def isAscending(val1, val2):
    return True if (val1 - val2 < 0) else False


main()