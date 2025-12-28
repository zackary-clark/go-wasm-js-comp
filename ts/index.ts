globalThis.addJS = addJS;

function addJS() {
    const form: HTMLFormElement = globalThis.document.getElementById("add-form");
    const a = form.elements["a"].valueAsNumber;
    const b = form.elements["b"].valueAsNumber;
    const c = a + b;
    const output = form.elements["js-out"];
    output.value = c;
}
