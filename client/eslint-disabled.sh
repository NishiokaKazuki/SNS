# client/proto/*.jsの先頭行に /* eslint-disable */ を追加
find client/src/proto/*.js | xargs gsed -i '1i /* eslint-disable */'