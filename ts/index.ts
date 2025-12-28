globalThis.multJS = multJS;
globalThis.addJS = addJS;

function multJS() {
    const form: HTMLFormElement = globalThis.document.getElementById("mult-form");
    const a = form.elements["a"].valueAsNumber;
    const b = form.elements["b"].valueAsNumber;
    const c = a * b;
    const output = form.elements["js-out"];
    output.value = c;
}

function addJS() {
    const form: HTMLFormElement = globalThis.document.getElementById("add-form");
    const a = form.elements["a"].valueAsNumber;
    const b = form.elements["b"].valueAsNumber;
    const c = a + b;
    const output = form.elements["js-out"];
    output.value = c;
}
