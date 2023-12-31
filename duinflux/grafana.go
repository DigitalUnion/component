package duinflux

var grafanaTemplate = `
{
    "dashboard":{
        "id":null,
        "title":"$title",
        "uid":"$title",
        "timezone":"",
        "editable":true,
        "gnetId":null,
        "graphTooltip":0,
        "links":[

        ],
        "panels":[
            {
                "datasource":null,
                "fieldConfig":{
                    "defaults":{
                        "color":{
                            "mode":"thresholds"
                        },
                        "custom":{
                            "align":"center",
                            "displayMode":"color-text",
                            "filterable":false
                        },
                        "mappings":[

                        ],
                        "thresholds":{
                            "mode":"absolute",
                            "steps":[
                                {
                                    "color":"green",
                                    "value":null
                                },
                                {
                                    "color":"red",
                                    "value":80
                                }
                            ]
                        }
                    },
                    "overrides":[
                        {
                            "matcher":{
                                "id":"byName",
                                "options":"_value"
                            },
                            "properties":[
                                {
                                    "id":"unit",
                                    "value":"s"
                                },
                                {
                                    "id":"displayName",
                                    "value":"运行时间"
                                }
                            ]
                        },
                        {
                            "matcher":{
                                "id":"byName",
                                "options":"ip"
                            },
                            "properties":[
                                {
                                    "id":"displayName",
                                    "value":"IP"
                                }
                            ]
                        }
                    ]
                },
                "gridPos":{
                    "h":9,
                    "w":12,
                    "x":0,
                    "y":0
                },
                "id":8,
                "options":{
                    "frameIndex":0,
                    "showHeader":true
                },
                "pluginVersion":"8.1.1",
                "targets":[
                    {
                        "query":"from(bucket: \"${bucket}\")\n  |> range(start: -12s, stop: now())\n  |> filter(fn: (r) => r[\"_measurement\"] == \"uptime\")\n  |> filter(fn: (r) => r[\"_field\"] == \"seconds\")\n  |> group(columns: [\"_measurement\", \"_time\",\"_field\"])\n  |> aggregateWindow(every:12s, fn: last, createEmpty: false)\n  |> sort(columns: [\"ip\"])\n  |> yield(name: \"last\")",
                        "refId":"A"
                    }
                ],
                "title":"节点运行时间",
                "type":"table"
            },
            {
                "datasource":null,
                "fieldConfig":{
                    "defaults":{
                        "color":{
                            "mode":"palette-classic"
                        },
                        "custom":{
                            "axisLabel":"",
                            "axisPlacement":"auto",
                            "barAlignment":0,
                            "drawStyle":"line",
                            "fillOpacity":10,
                            "gradientMode":"opacity",
                            "hideFrom":{
                                "legend":false,
                                "tooltip":false,
                                "viz":false
                            },
                            "lineInterpolation":"smooth",
                            "lineStyle":{
                                "fill":"solid"
                            },
                            "lineWidth":1,
                            "pointSize":2,
                            "scaleDistribution":{
                                "type":"linear"
                            },
                            "showPoints":"auto",
                            "spanNulls":true,
                            "stacking":{
                                "group":"A",
                                "mode":"normal"
                            },
                            "thresholdsStyle":{
                                "mode":"off"
                            }
                        },
                        "mappings":[

                        ],
                        "min":0,
                        "thresholds":{
                            "mode":"absolute",
                            "steps":[
                                {
                                    "color":"green",
                                    "value":null
                                },
                                {
                                    "color":"red",
                                    "value":80
                                }
                            ]
                        },
                        "unit":"short"
                    },
                    "overrides":[

                    ]
                },
                "gridPos":{
                    "h":9,
                    "w":12,
                    "x":12,
                    "y":0
                },
                "id":2,
                "options":{
                    "legend":{
                        "calcs":[

                        ],
                        "displayMode":"list",
                        "placement":"bottom"
                    },
                    "tooltip":{
                        "mode":"multi"
                    }
                },
                "pluginVersion":"8.1.1",
                "targets":[
                    {
                        "query":"from(bucket: \"${bucket}\")\n  |> range(start: v.timeRangeStart, stop: v.timeRangeStop)\n  |> filter(fn: (r) => r[\"_measurement\"] =~ /gin_200_POST/)\n  |> filter(fn: (r) => r[\"_field\"] == \"qps\")\n  |> aggregateWindow(every: ${interval}, fn: mean, createEmpty: false)\n  |> yield(name: \"mean\")",
                        "refId":"A"
                    }
                ],
                "timeFrom":null,
                "timeShift":null,
                "title":"QPS",
                "type":"timeseries"
            },
            {
                "datasource":null,
                "fieldConfig":{
                    "defaults":{
                        "color":{
                            "mode":"palette-classic"
                        },
                        "custom":{
                            "axisLabel":"",
                            "axisPlacement":"auto",
                            "barAlignment":0,
                            "drawStyle":"line",
                            "fillOpacity":10,
                            "gradientMode":"opacity",
                            "hideFrom":{
                                "legend":false,
                                "tooltip":false,
                                "viz":false
                            },
                            "lineInterpolation":"smooth",
                            "lineWidth":1,
                            "pointSize":2,
                            "scaleDistribution":{
                                "type":"linear"
                            },
                            "showPoints":"auto",
                            "spanNulls":true,
                            "stacking":{
                                "group":"A",
                                "mode":"none"
                            },
                            "thresholdsStyle":{
                                "mode":"off"
                            }
                        },
                        "mappings":[

                        ],
                        "min":0,
                        "thresholds":{
                            "mode":"absolute",
                            "steps":[
                                {
                                    "color":"green",
                                    "value":null
                                },
                                {
                                    "color":"red",
                                    "value":80
                                }
                            ]
                        },
                        "unit":"µs"
                    },
                    "overrides":[

                    ]
                },
                "gridPos":{
                    "h":9,
                    "w":12,
                    "x":0,
                    "y":9
                },
                "id":4,
                "options":{
                    "legend":{
                        "calcs":[

                        ],
                        "displayMode":"list",
                        "placement":"bottom"
                    },
                    "tooltip":{
                        "mode":"multi"
                    }
                },
                "pluginVersion":"8.1.1",
                "targets":[
                    {
                        "query":"from(bucket: \"${bucket}\")\n  |> range(start: v.timeRangeStart, stop: v.timeRangeStop)\n  |> filter(fn: (r) => r[\"_measurement\"] =~ /gin_200_POST_/)\n  |> filter(fn: (r) => r[\"_field\"] == \"cost\")\n  |> aggregateWindow(every: ${interval}, fn: mean, createEmpty: false)\n  |> yield(name: \"mean\")",
                        "refId":"A"
                    }
                ],
                "timeFrom":null,
                "timeShift":null,
                "title":"节点平均耗时",
                "type":"timeseries"
            },
            {
                "datasource":null,
                "fieldConfig":{
                    "defaults":{
                        "color":{
                            "mode":"palette-classic"
                        },
                        "custom":{
                            "axisLabel":"",
                            "axisPlacement":"auto",
                            "barAlignment":0,
                            "drawStyle":"line",
                            "fillOpacity":10,
                            "gradientMode":"opacity",
                            "hideFrom":{
                                "legend":false,
                                "tooltip":false,
                                "viz":false
                            },
                            "lineInterpolation":"smooth",
                            "lineStyle":{
                                "fill":"solid"
                            },
                            "lineWidth":1,
                            "pointSize":2,
                            "scaleDistribution":{
                                "type":"linear"
                            },
                            "showPoints":"auto",
                            "spanNulls":false,
                            "stacking":{
                                "group":"A",
                                "mode":"none"
                            },
                            "thresholdsStyle":{
                                "mode":"off"
                            }
                        },
                        "mappings":[

                        ],
                        "min":0,
                        "thresholds":{
                            "mode":"absolute",
                            "steps":[
                                {
                                    "color":"green",
                                    "value":null
                                },
                                {
                                    "color":"red",
                                    "value":80
                                }
                            ]
                        },
                        "unit":"µs"
                    },
                    "overrides":[
                        {
                            "matcher":{
                                "id":"byFrameRefID",
                                "options":"A"
                            },
                            "properties":[
                                {
                                    "id":"displayName",
                                    "value":"99%"
                                }
                            ]
                        },
                        {
                            "matcher":{
                                "id":"byFrameRefID",
                                "options":"B"
                            },
                            "properties":[
                                {
                                    "id":"displayName",
                                    "value":"95%"
                                }
                            ]
                        },
                        {
                            "matcher":{
                                "id":"byFrameRefID",
                                "options":"C"
                            },
                            "properties":[
                                {
                                    "id":"displayName",
                                    "value":"90%"
                                }
                            ]
                        }
                    ]
                },
                "gridPos":{
                    "h":9,
                    "w":12,
                    "x":12,
                    "y":9
                },
                "id":6,
                "options":{
                    "legend":{
                        "calcs":[

                        ],
                        "displayMode":"list",
                        "placement":"bottom"
                    },
                    "tooltip":{
                        "mode":"multi"
                    }
                },
                "targets":[
                    {
                        "query":"from(bucket: \"${bucket}\")\n  |> range(start: v.timeRangeStart, stop: v.timeRangeStop)\n  |> filter(fn: (r) => r[\"_measurement\"] =~ /gin_200_POST_/)\n  |> filter(fn: (r) => r[\"_field\"] == \"cost\")\n  |> group(columns:[\"_measurement\",\"_field\"])\n   |> aggregateWindow(every: ${interval}, fn: (tables=<-, column) => tables\n            |> quantile(q: 0.99, method: \"exact_selector\")\n  )\n  |> yield(name: \"mean\")",
                        "refId":"A"
                    },
                    {
                        "hide":false,
                        "query":"from(bucket: \"${bucket}\")\n  |> range(start: v.timeRangeStart, stop: v.timeRangeStop)\n  |> filter(fn: (r) => r[\"_measurement\"] =~ /gin_200_POST_/)\n  |> filter(fn: (r) => r[\"_field\"] == \"cost\")\n  |> group(columns:[\"_measurement\",\"_field\"])\n   |> aggregateWindow(every: ${interval}, fn: (tables=<-, column) => tables\n            |> quantile(q: 0.95, method: \"exact_selector\")\n  )\n  |> yield(name: \"mean\")",
                        "refId":"B"
                    },
                    {
                        "hide":false,
                        "query":"from(bucket: \"${bucket}\")\n  |> range(start: v.timeRangeStart, stop: v.timeRangeStop)\n  |> filter(fn: (r) => r[\"_measurement\"] =~ /gin_200_POST_/)\n  |> filter(fn: (r) => r[\"_field\"] == \"cost\")\n  |> group(columns:[\"_measurement\",\"_field\"])\n   |> aggregateWindow(every: ${interval}, fn: (tables=<-, column) => tables\n            |> quantile(q: 0.90, method: \"exact_selector\")\n  )\n  |> yield(name: \"mean\")",
                        "refId":"C"
                    }
                ],
                "title":"耗时百分比",
                "type":"timeseries"
            }
        ],
        "style":"dark",
        "tags":[

        ],
        "templating":{
            "list":[
                {
                    "auto":false,
                    "auto_count":30,
                    "auto_min":"10s",
                    "current":{
                        "selected":false,
                        "text":"1m",
                        "value":"1m"
                    },
                    "description":null,
                    "error":null,
                    "hide":0,
                    "label":null,
                    "name":"interval",
                    "options":[
                        {
                            "selected":true,
                            "text":"1m",
                            "value":"1m"
                        },
                        {
                            "selected":false,
                            "text":"10m",
                            "value":"10m"
                        },
                        {
                            "selected":false,
                            "text":"30m",
                            "value":"30m"
                        },
                        {
                            "selected":false,
                            "text":"1h",
                            "value":"1h"
                        },
                        {
                            "selected":false,
                            "text":"6h",
                            "value":"6h"
                        },
                        {
                            "selected":false,
                            "text":"12h",
                            "value":"12h"
                        },
                        {
                            "selected":false,
                            "text":"1d",
                            "value":"1d"
                        },
                        {
                            "selected":false,
                            "text":"7d",
                            "value":"7d"
                        },
                        {
                            "selected":false,
                            "text":"14d",
                            "value":"14d"
                        },
                        {
                            "selected":false,
                            "text":"30d",
                            "value":"30d"
                        }
                    ],
                    "query":"1m,10m,30m,1h,6h,12h,1d,7d,14d,30d",
                    "queryValue":"",
                    "refresh":2,
                    "skipUrlSync":false,
                    "type":"interval"
                },
                {
                    "description":null,
                    "error":null,
                    "hide":2,
                    "label":null,
                    "name":"bucket",
                    "query":"$bucket",
                    "skipUrlSync":false,
                    "type":"constant"
                }
            ]
        },
        "schemaVersion":30,
        "version":18
    },
    "overwrite":true
}`
