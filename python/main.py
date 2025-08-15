with open(0) as f:
    in_pre = False
    in_list = False
    while line := f.readline():
        if in_pre:
            if line == '```':
                print('</pre>')
                in_pre = False
            else:
                print(line)
            continue

        if in_list:
            if not line.startswith('* '):
                in_list = False
                print('</ul>')

        l = line.strip()
        if not l: print('<br>')
        elif l.startswith('### '): print(f'<h3>{l[4:]}</h3>')
        elif l.startswith('## '): print(f'<h2>{l[3:]}</h2>')
        elif l.startswith('# '): print(f'<h1>{l[2:]}</h1>')
        elif l.startswith('=>'):
            url, *text = l[2:].strip().split(' ', 1)
            print(f'<a href="{url}">{text[0] if len(text) else ""}</a>')
        elif  l.startswith('>'):
            print(f'<blockquote>{l[1:]}</blockquote>')
        elif l.startswith('* '):
            if not in_list:
                in_list = True
                print('<ul>')
            print(f'\t<li>{l[2:]}</li>')
        elif l.startswith('```'):
            in_pre = True
            print('<pre>')
        else: print(f'<p>{line.strip('\n\r')}</p>')

    if in_list: print('</ul>')
    if in_pre: print('</pre>')
