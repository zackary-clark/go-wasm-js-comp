export function add() {
    const form = globalThis.document.getElementById("add-form") as HTMLFormElement;
    const a = form.elements["a"].valueAsNumber;
    const b = form.elements["b"].valueAsNumber;
    const c = a + b;
    const output = form.elements["js-out"];
    output.value = c;
}
