rm -rf www
rm README-API.md
apidoc -o www
apidoc-markdown2 -p www -o README-API.md
