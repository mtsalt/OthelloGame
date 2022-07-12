class Othello extends GuiDrawer{

    constructor(canvas_id, windowSize){
        super(canvas_id, windowSize, windowSize);
        setMesh();
    }

    drawBoard(boardArr) {
        let x = 15;
        let y = 15;
        let size = 80;
        let outlineColor = COLOR_BLACK;
        let lineWidth = 5;

        for(let i=0; i<8; i++) {
            for(let j=0; j<8; j++){
                let fillColor = this.getBoardColor(boardArr[j][i]);
                super.square(x+size*i, y+size*j, size, fillColor, outlineColor, lineWidth);
            }
        }
    }

    drawStone(boardArr) {
        let x = 55;
        let y = 55;
        let radius = 30;
        let size = 80;
        
        for(let i=0; i<8; i++) {
            for(let j=0; j<8; j++) {
                if(boardArr[i][j] != 0){
                    let fillColor = this.getStoneColor2(boardArr[i][j]);
                    super.circle(x+size*j, y+size*i, radius, fillColor);
                }
            }
        }
    }

    update(boardArr) {
        this.drawBoard(boardArr);
        this.drawStone(boardArr);
    }

    getBoardColor(boardState) {
        switch (boardState) {
            case BOARD_STATE_AVAILABLE:
                return BOARD_CELL_COLOR_AVAILABLE;
            default:
                return BOARD_CELL_COLOR_EMPTY;
        }
    }

    getStoneColor(boardState) {
        switch (boardState) {
            case 0:
                return "rgba(34,139,34,1)";
            case 1:
                return "white";
            case 2:
                return "black"
            case 3:
                return "rgba(0,255,127,1)";
            default:
                return "rgba(34,139,34,1)";
        }
    }

    getStoneColor2(boardState) {
        switch (boardState) {
            case BOARD_STATE_EMPTY:
                return STONE_EMPTY;
            case BOARD_STATE_WHITE:
                return STONE_WHITE;
            case BOARD_STATE_BLACK:
                return STONE_BLACK;
            case BOARD_STATE_AVAILABLE:
                return STONE_AVAILABLE;
            default:
                return STONE_EMPTY;
        }
    }
}

function setMesh() {

    let mesh = document.getElementById("mesh");

    // div (meshの中身を空にしておく必要あり。呼ぶ度にどんどん増えていくから） 

    for(let i=0; i<8; i++) {
        for(let j=0; j<8; j++) {
            let id = "mesh_" + i.toString() + "-" +  j.toString();
            let elem = document.createElement("div");
            elem.id = id;
            elem.style.display = "inline-block";
            elem.style.margin = "1px 4px 0px 1px";
            elem.onclick = meshClickEvent;
            elem.style.width = "75px";
            elem.style.height = "75px"
            mesh.appendChild(elem);

            // let elem2 = document.getElementById(id);
            // elem2.onclick = meshClickEvent;
        }
    }
}

function meshClickEvent(e) {

    let meshId = e.target.id;
    axisData = extractAxisFromMeshId(meshId);

    // このイベントでは、座標取得＆辞書形式へ変更、だけを行うようにする
    // 　→　どうにかして、①送信可能かどうかboard配列の内容から判断　②良ければWebSocketで送信
    //　　　をする必要がある
    wsSend(axisData);
}

function test2() {
    socket.send("hello");
    wsSend(1,2);
}

function extractAxisFromMeshId(meshId) {
    let axisData = meshId.split("_")[1].split("-");
    return {x: axisData[0], y: axisData[1]};
}

class GameInfo extends GuiDrawer{

    constructor(canvas_id, windowSizeX, windowSizeY){
        super(canvas_id, windowSizeX, windowSizeY);
    }
}