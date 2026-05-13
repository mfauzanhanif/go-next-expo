import zipfile
with zipfile.ZipFile("/tmp/echo-v5/v5.1.1.zip", 'r') as zip_ref:
    zip_ref.extractall("/tmp/echo-v5")
