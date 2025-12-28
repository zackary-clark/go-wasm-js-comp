console.log("from typescript")

function add() {
    safelyCall(addJS);
    safelyCall(addWASM);
}

function safelyCall(callback: Function) {
    const timerLabel = callback?.name || "safelyCall";
    let out: any;
    console.time(timerLabel);
    try {
        out = callback();
    } catch (e) {
        console.error("error thrown by %s: ", timerLabel, e);
    } finally {
        console.timeEnd(timerLabel);
    }
    return out;
}

globalThis.add = add;

function addWASM() {
    globalThis.addWASM()
}

function addJS() {
    const a = getNumberFromInput("a");
    const b = getNumberFromInput("b");
    console.log("a: ", a);
    console.log("b: ", b);
    const c = a + b;
    outputToInput("out", c);
    console.log("c: ", c);
}

function outputToInput(className: string, output: number) {
    const outputElement = getElementByClassName("js-output")?.getElementsByClassName(className)[0];
    outputElement.value = output;
}

function getNumberFromInput(className: string) {
    const value = getElementByClassName(className)?.value;
    return parseFloat(value);
}

function getElementByClassName(className: string) {
    return getAddArticle()?.getElementsByClassName(className)[0];
}

function getAddArticle() {
    return globalThis.document.getElementById("add")
}
