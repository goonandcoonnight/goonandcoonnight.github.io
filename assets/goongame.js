const goWasm = new Go()

WebAssembly.instantiateStreaming(fetch("goongame.wasm"), goWasm.importObject).then((result) => {
    goWasm.run(result.instance)
    document.getElementById("print-html").addEventListener("click", () => {
        document.body.innerHTML += printHtml()
    })
})