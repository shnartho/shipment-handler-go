function order() {
    document.getElementById("packForm").action = "/order";
    document.getElementById("packForm").submit();
}

function add() {
    document.getElementById("packForm").action = "/add";
    document.getElementById("packForm").submit();
}

function update() {
    document.getElementById("packForm").action = "/update";
    document.getElementById("packForm").submit();
}

function remove() {
    document.getElementById("packForm").action = "/remove";
    document.getElementById("packForm").submit();
}

function orderPack() {
    document.getElementById("packForm").action = "/orderpack";
    document.getElementById("packForm").submit();
}

function addPack() {
    document.getElementById("packForm").action = "/addpack";
    document.getElementById("packForm").submit();
}

function updatePack() {
    document.getElementById("packForm").action = "/updatepack";
    document.getElementById("packForm").submit();
}

function removePack() {
    document.getElementById("packForm").action = "/removepack";
    document.getElementById("packForm").submit();
}