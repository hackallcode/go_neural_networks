function Post(url, data, callback) {
    $.ajax({
        type: 'POST',
        url: url,
        data: JSON.stringify(data),
        dataType: 'json',
    })
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

function Delete(url, data, callback) {
    $.ajax({
        type: 'DELETE',
        url: url,
        data: JSON.stringify(data),
        dataType: 'json',
    })
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

function ApiAddArea(callback) {
    Post("/api/area", {}, callback)
}

function ApiAddPoints(areaId, points, callback) {
    Post("/api/point", {id: areaId, points: points}, callback)
}

function ApiAddClusters(areaId, clusters, callback) {
    Post("/api/cluster", {id: areaId, clusters: clusters}, callback)
}

function ApiTrain(areaId, distanceId, byStep, maxAge, callback) {
    Post("/api/train", {id: areaId, dist_id: distanceId, by_step: byStep, max_age: maxAge}, callback)
}

function ApiClearArea(areaId, callback) {
    Delete("/api/area", {id: areaId}, callback)
}
