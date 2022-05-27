// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.

require('froala-editor/js/froala_editor.pkgd.min.js')
require('froala-editor/js/plugins/image.min.js')
require('froala-editor/js/plugins/table.min.js')
require('froala-editor/js/plugins/colors.min.js')
require('froala-editor/js/plugins/draggable.min.js')
require('froala-editor/js/plugins/align.min.js')
require('froala-editor/js/plugins/font_size.min.js')
require('froala-editor/js/plugins/lists.min.js')
require('froala-editor/js/plugins/link.min.js')
require('froala-editor/js/plugins/line_height.min.js')
require('froala-editor/js/plugins/line_breaker.min.js')
require('froala-editor/js/plugins/font_family.min.js')
require('froala-editor/js/plugins/entities.min.js')
require('froala-editor/js/plugins/paragraph_format.min.js')
require('froala-editor/js/plugins/paragraph_style.min.js')
require('froala-editor/js/plugins/print.min.js')
require('froala-editor/js/plugins/quote.min.js')
require('froala-editor/js/plugins/video.min.js')
require('froala-editor/js/plugins/word_paste.min.js')
require('froala-editor/js/plugins/url.min.js')
require('froala-editor/js/plugins/special_characters.min.js')
require('froala-editor/js/plugins/emoticons.min.js')
require('froala-editor/js/plugins/code_beautifier.min.js')

// require('froala-editor/js/plugins/file.min.js')

require('froala-editor/css/froala_editor.pkgd.min.css')
require('froala-editor/css/froala_style.min.css')
require('froala-editor/css/plugins/image.min.css')
require('froala-editor/css/plugins/table.min.css')
require('froala-editor/css/plugins/colors.min.css')
require('froala-editor/css/plugins/draggable.min.css')
require('froala-editor/css/plugins/video.min.css')
require('froala-editor/css/plugins/special_characters.min.css')
require('froala-editor/css/plugins/line_breaker.min.css')
require('froala-editor/css/plugins/emoticons.min.css')

import Vue from 'vue'
import VueFroala from 'vue-froala-wysiwyg'
import Axios from 'axios'
import axios from 'axios'
import App from './App'
import router from './router'
import store from './store'

Vue.config.productionTip = false
Vue.prototype.$http = Axios
Vue.use(VueFroala)

// axios.defaults.baseURL = 'http://localhost:8000';

// Important: If axios is used with multiple domains, the AUTH_TOKEN will be sent to all of them.
// See below for an example using Custom instance defaults instead.
// axios.defaults.headers.common['Authorization'] = AUTH_TOKEN;

axios.defaults.headers.post['Content-Type'] = "application/json"

// instance.defaults.headers.common['Authorization'] = AUTH_TOKEN;

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
