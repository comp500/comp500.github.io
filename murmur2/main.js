document.getElementById("result").innerText = "Preparing...";

if (!WebAssembly.instantiateStreaming) { // polyfill
	WebAssembly.instantiateStreaming = async (resp, importObject) => {
		const source = await (await resp).arrayBuffer();
		return await WebAssembly.instantiate(source, importObject);
	};
}

const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
	go.run(result.instance);

	let fileReader = new FileReader();
	fileReader.onloadend = (e) => {
		let arr = new Uint8Array(fileReader.result);
		document.getElementById("result").innerText = computeHash(arr);
	};

	let inputFile = document.getElementById("inputfile");
	inputFile.addEventListener("change", (e) => {
		fileReader.readAsArrayBuffer(inputFile.files[0]);
	});

	document.getElementById("result").innerText = "Ready!";
});