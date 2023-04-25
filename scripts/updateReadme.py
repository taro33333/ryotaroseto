GENERATE_FILE_NAME = 'README.md'

# dedent: 各行の先頭の空白削除
docs_str = '![](./profile-3d-contrib/profile-south-season.svg)'

# READMEに書き込む
with open(GENERATE_FILE_NAME, 'w') as f:
    f.write(docs_str)
