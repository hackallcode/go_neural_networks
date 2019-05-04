let areaId = 0;

function Post(url, data, callback) {
    $.post(url, JSON.stringify(data))
        .done((answer) => {
            if (100 <= answer.status && answer.status < 200) {
                if (callback) {
                    callback(answer.data);
                }
            } else {
                console.error(answer.message);
            }
        })
        .fail((error) => {
            console.error(error);
        })
}


function ApiAddArea() {
    Post("/api/area", {}, (data) => {
        areaId = data.id;
        console.log("Area ID:", areaId);
    })
}

function ApiAddPoints(points, callback) {
    Post("/api/point", {id: areaId, points: points}, callback)
}

function ApiAddClusters(clusters, callback) {
    Post("/api/cluster", {id: areaId, clusters: clusters}, callback)
}

function ApiTrain(distanceId, byStep, maxAge, callback) {
    Post("/api/train", {id: areaId, dist_id: distanceId, by_step: byStep, max_age: maxAge}, callback)
}

ApiAddArea();