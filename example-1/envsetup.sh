app_setup() {
  echo "set up siot"
  (cd simpleiot &&
    . envsetup.sh &&
    siot_setup &&
    siot_build)
}

app_build() {
  go build -o my-app main.go
}
