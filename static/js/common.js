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
                    chunkSize: 2 * 1024 * 1024
                },
                attrs: {
                    accept: 'image/*'
                },
                previewInfo: ''
            }
        },
        mounted: function () {
            this.pingDevice();
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
                    vm.devShow = value.dev
                    vm.syncDevs = value.devs
                    vm.showSyncText = value.dev.sync_text
                    vm.showSyncFile = value.dev.sync_file
                    console.log(vm.syncDevs)
                })
            },
            uploadSuccess: function (rootFile, file, message) {
                console.log(rootFile, file.file, message)
                $.ajax("/dev/sync", {
                    method: "POST",
                    data: {
                        toMd5: this.choiceDev,
                        syncType: "file",
                        desc: file.file.name,

                    },
                    dataType: 'json'
                }).then(
                    function (val) {
                        console.log(val)
                    }
                )
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

            }

        },
        watch: {
            syncText: function (newval) {
                $.ajax("/dev/sync", {
                    method: "POST",
                    data: {
                        toMd5: this.choiceDev,
                        syncType: 'text',
                        content: newval,

                    },
                    dataType: 'json'
                }).then(
                    function (val) {
                        console.log(val)
                    }
                )
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


setInterval(function () {
    $.ajax(
        "/dev/ping",
        {method: "GET"}
    ).then(value => {
        vm.devShow = value.dev
        vm.syncDevs = value.devs
        vm.showSyncText = value.dev.sync_text
        vm.showSyncFile = value.dev.sync_file
        console.log(vm.showSyncText)
    })
}, 5000)

$(function () {
    var nua = navigator.userAgent
    var isAndroid = (nua.indexOf('Mozilla/5.0') > -1 && nua.indexOf('Android ') > -1 && nua.indexOf('AppleWebKit') > -1 && nua.indexOf('Chrome') === -1)
    if (isAndroid) {
        $('select.form-control').removeClass('form-control').css('width', '100%')
    }
})
