import os

def countlines(start, lines=0, header=True, begin_start=None):
    if header:
        print('{:>10} |{:>10} | {:<20}'.format('ADDED', 'TOTAL', 'FILE'))
        print('{:->11}|{:->11}|{:->20}'.format('', '', ''))
    for file in os.listdir(start):
        file = os.path.join(start, file)
        if os.path.isfile(file):
            if file.endswith('.go') and not file.endswith('.pb.go'):
                with open(file, 'r') as f:
                    newlines = f.readlines()
                    newlines = len(newlines)
                    lines += newlines
                    if begin_start is not None:
                        reldir_of_thing = '.' + file.replace(begin_start, '')
                    else:
                        reldir_of_thing = '.' + file.replace(start, '')
                    print('{:>10} |{:>10} | {:<20}'.format(
                            newlines, lines, reldir_of_thing))
    for file in os.listdir(start):
        file = os.path.join(start, file)
        if os.path.isdir(file):
            lines = countlines(file, lines, header=False, begin_start=start)
    return lines

print(countlines('.'))