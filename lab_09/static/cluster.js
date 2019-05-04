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

let points = [];
let clusters = [];

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

function train(byStep) {
    ApiAddPoints(points, () => {
        ApiAddClusters(clusters, () => {
            ApiTrain($('distId').val(), byStep, $('maxAge').val(), (data) => {
                ClearCanvas();
                points = [];
                clusters = [];
                lastColor = 0;

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
