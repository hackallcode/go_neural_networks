const pointsColor = '#000000';

let colors = [
    '#F7E017',
    '#ED6E00',
    '#E60094',
    '#7800b9',
    '#2905A1',
    '#008CCC',
    '#00B394',
    '#00AB59',
    '#F54029',
    '#FC8CA1',
    '#E66BC2',
    '#B878BF',
    '#73B5E0',
];
let lastColor = 0;

let areaId = -1;
let points = [];
let clusters = [];

function clearAll() {
    ClearCanvas();
    points = [];
    clusters = [];
    lastColor = 0;
}

function printPoint(ctx, x, y, color) {
    ctx.fillStyle = color;
    ctx.fillRect(x, y, 1, 1);
}

function addPoint(ctx, x, y, color) {
    printPoint(ctx, x, y, color);
    points.push({x: x, y: y});
}

function printCluster(ctx, x, y, color) {
    ctx.fillStyle = color;
    ctx.fillRect(x - 1, y, 1, 1);
    ctx.fillRect(x, y - 1, 1, 1);
    ctx.fillRect(x + 1, y, 1, 1);
    ctx.fillRect(x, y + 1, 1, 1);
    ctx.fillRect(x, y, 1, 1);
}

function addCluster(ctx, x, y, color) {
    printCluster(ctx, x, y, color);
    clusters.push({x: x, y: y});
}

function train(distId, byStep, maxAge) {
    ApiAddPoints(areaId, points, () => {
        ApiAddClusters(areaId, clusters, () => {
            ApiTrain(areaId, distId, byStep, maxAge, (data) => {
                clearAll();

                data.clusters.forEach((cluster) => {
                    let color = colors[lastColor];
                    lastColor++;
                    lastColor %= colors.length;

                    printCluster(ctx, cluster.x, cluster.y, color);
                    cluster.points.forEach((point) => {
                        printPoint(ctx, point.x, point.y, color);
                    });
                });
            })
        });
    });
}

function leftClick(event) {
    event.preventDefault();
    addPoint(ctx, Math.trunc(event.offsetX / SCALE), Math.trunc(event.offsetY / SCALE), pointsColor);
}

function rightClick(event) {
    event.preventDefault();
    addCluster(ctx, Math.trunc(event.offsetX / SCALE), Math.trunc(event.offsetY / SCALE), colors[lastColor]);
    lastColor++;
    lastColor %= colors.length;
}

// Main code
let SCALE = 10;
let canvas = CanvasInit('canvas', SCALE);
let ctx = canvas.getContext('2d');
canvas.addEventListener('click', leftClick);
canvas.addEventListener('contextmenu', rightClick);

ApiAddArea((data) => {
    areaId = data.id;
    console.log("Area ID:", areaId);
});

// Test data
function addTestData() {
    points = [
        {"x": 13, "y": 11},
        {"x": 31, "y": 13},
        {"x": 13, "y": 28},
        {"x": 30, "y": 36},
        {"x": 20, "y": 18},
        {"x": 42, "y": 22},
        {"x": 18, "y": 39},
        {"x": 26, "y": 27},
        {"x": 11, "y": 22},
        {"x": 23, "y": 8},
        {"x": 45, "y": 39},
        {"x": 11, "y": 37},
        {"x": 47, "y": 14},
        {"x": 21, "y": 20},
        {"x": 39, "y": 32},
        {"x": 34, "y": 23},
        {"x": 6, "y": 6},
        {"x": 41, "y": 9},
        {"x": 18, "y": 7},
        {"x": 29, "y": 21},
        {"x": 4, "y": 21},
        {"x": 32, "y": 9},
        {"x": 6, "y": 12},
        {"x": 25, "y": 14},
        {"x": 6, "y": 28},
        {"x": 37, "y": 36},
        {"x": 15, "y": 45},
        {"x": 39, "y": 20},
        {"x": 16, "y": 31},
        {"x": 30, "y": 30},
        {"x": 19, "y": 26},
        {"x": 15, "y": 18},
        {"x": 29, "y": 7},
        {"x": 13, "y": 6}
    ];
    clusters = [
        {"x": 10, "y": 14},
        {"x": 36, "y": 17},
        {"x": 23, "y": 33}
    ];
    train(1, false, 100);
}
