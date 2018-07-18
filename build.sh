GOPATH="$(pwd)"
export GOPATH=$GOPATH

if [[ ! -d "$GOPATH/bin/" ]]; then
    mkdir -p "$GOPATH/bin/"
fi

# Clean binary output
rm -rf $GOPATH/bin/*.*

go build -o $GOPATH/bin/go_eval_example go_eval_example

if [ $? -eq 0 ]; then
    cp $GOPATH/src/go_eval_example/config.toml $GOPATH/bin/config.toml
    echo "SUCCESS"
else
    echo "FAILED"
fi
