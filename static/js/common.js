window.vm = new Vue(
    {
        el: "#app",
        data() {
            return {
                syncDevs: [],
                devShow: {
                    show: 'waiting...'
                },
                choiceDev: '',
                syncText: '',
                showSyncText: [],
                showSyncFile: [],
                UploadOptions: {
                    // https://github.com/simple-uploader/Uploader/tree/develop/samples/Node.js
                    target: '/dev/upload',
                    testChunks: false,
                    chunkSize: 10 * 1024 * 1024
                },
                attrs: {
                    accept: 'image/*'
                },
                previewInfo: '',
                ws: null, // websocket
                wsMsgType: {
                    ping: 'ping',
                    pingGroup: 'pingGroup',
                    sync: 'sync',
                },

            }
        },
        mounted: function () {
            this.pingDevice();

            // 6min发起一次心跳检测,ws断开后1min发起一次重试
            setInterval(
                () => {
                    if (this.ws !== null) {
                        this.ws.send(JSON.stringify({type: this.wsMsgType.ping}))
                    }
                },
                1000 * 60 * 6
            )
            setInterval(
                () => {
                    if (this.ws == null) {
                        this.pingDevice()
                    }
                },
                1000 * 60
            )

        },
        methods: {
            choiceDevClick: function (md) {
                this.choiceDev = md
            },
            pingDevice: function () {
                $.ajax(
                    "/dev/ping",
                    {method: "GET"}
                ).then(value => {
                    this.InitWebSocketSvr()
                })
            },
            uploadSuccess: function (rootFile, file, message) {
                console.log(rootFile, file.file, message)
                this.ws.send(JSON.stringify({
                    type: this.wsMsgType.sync,
                    content: JSON.stringify({
                        toMd5: this.choiceDev,
                        syncType: "file",
                        desc: file.file.name,
                    })
                }))

            },
            generateFileIcon: function (ext) {
                let dfHtml = '<i class="bi-file-earmark"></i>'
                if (
                    [
                        '.js',
                        '.go',
                        '.py',
                        ".mp4",
                        ".av",
                        ".png",
                        ".jpg",
                        ".jpeg",
                        ".webp",
                    ].indexOf(ext) !== -1
                ) {
                    return `<i class="bi-filetype-${ext.replace('.', '')}"></i>`
                }

            },
            previewFileClick: function (ext, path) {
                let otherHtml = `<a target="_blank" href="${path}">点击查看</a>`
                let imgHtml = `<img style="height: 550px" src="${path}" class="figure-img img-thumbnail img-fluid rounded align-content-center" alt="...">`
                let mp4Html = `<video height="200" width="320" controls autoplay="true" src="${path}"></video>`
                switch (ext) {
                    case ".png":
                        vm.previewInfo = imgHtml
                        break
                    case ".jpg":
                        vm.previewInfo = imgHtml
                        break
                    case '.jpeg':
                        vm.previewInfo = imgHtml
                        break
                    case '.webp':
                        vm.previewInfo = imgHtml
                        break
                    case '.mp4':
                        vm.previewInfo = mp4Html
                        break
                    case '.av':
                        vm.previewInfo = mp4Html
                        break
                    default:
                        vm.previewInfo = otherHtml
                        break
                }
                $("#syncPreviewModal").modal('show')

            },
            InitWebSocketSvr: function () {
                this.ws = new WebSocket(`ws://${window.location.host}/dev/ws/1`);
                this.ws.onopen = () => {
                    console.log('ws连接状态：' + this.ws.readyState);
                    //连接成功 ping group
                    this.ws.send(JSON.stringify({type: this.wsMsgType.pingGroup}))
                }
                this.ws.onmessage = (msg) => {
                    const reader = new FileReader()
                    reader.readAsText(msg.data)
                    reader.onload = () => {
                        let val = JSON.parse(reader.result)
                        console.log('接收到来自服务器的消息：', Date.now(), val);
                        switch (val.type) {
                            case "ping": {
                                this.ws.send(JSON.stringify({type: this.wsMsgType.ping}))
                                break
                            }
                            case "sync": {
                                vm.devShow = val.val.dev
                                vm.syncDevs = val.val.devs
                                vm.showSyncText = val.val.dev.sync_text
                                vm.showSyncFile = val.val.dev.sync_file
                                break
                            }
                        }

                    }

                }
                this.ws.onerror = (errmsg) => {
                    console.log("ws err", errmsg)
                    this.ws = null
                }
                this.ws.onclose = () => {
                    console.log("ws close")
                    this.ws = null
                }
            }

        },
        watch: {
            syncText: function (newval) {
                this.ws.send(JSON.stringify({
                    type: this.wsMsgType.sync,
                    content: JSON.stringify({
                        toMd5: this.choiceDev,
                        syncType: 'text',
                        content: newval,
                    })
                }))
            },
            choiceDev: function () {
                console.log(this.choiceDev)
                if (this.choiceDev.length > 0) {
                    $("[id=uploadCard]").attr('style', 'display: block;')
                    $("[id=syncTextVal]").attr('style', 'display: block;')
                } else {
                    $("[id=uploadCard]").attr('style', 'display: none;')
                    $("[id=syncTextVal]").attr('style', 'display: none;')
                }
            }

        },

    }
)


// setInterval(function () {
//     $.ajax(
//         "/dev/ping",
//         {method: "GET"}
//     ).then(value => {
//         vm.devShow = value.dev
//         vm.syncDevs = value.devs
//         vm.showSyncText = value.dev.sync_text
//         vm.showSyncFile = value.dev.sync_file
//         console.log(vm.showSyncText)
//     })
// }, 5000)

$(function () {
    var nua = navigator.userAgent
    var isAndroid = (nua.indexOf('Mozilla/5.0') > -1 && nua.indexOf('Android ') > -1 && nua.indexOf('AppleWebKit') > -1 && nua.indexOf('Chrome') === -1)
    if (isAndroid) {
        $('select.form-control').removeClass('form-control').css('width', '100%')
    }
})
