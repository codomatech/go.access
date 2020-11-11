/*
 * TODO
 * - no need for mvvm library, all rendering is done once on startup
 * - find a way to update data of an apex chart
 * - for now, always use line charts
 * - every section (i.e. analysis) will specify:
 *    . x access names
 *    . data as array of records
 *    . which field is used for aggregation (summed to produce y)
 * - on zooming the chart, only the data zoomed will be shown in a table, along
 *   with useful insights (min, max, average & median for numeric fields, most and least
 *   frequent for string fields)
 * - perhaps do pie charts according to selection
 */
import ApexCharts from 'apexcharts'
import m from 'mithril'

let $state = {}

const AnalysisSection = {

    view: function(vnode) {
        const r = vnode.attrs.data
        const id = vnode.attrs.id
        if (!this.chart)
            setTimeout(() => m.redraw(), 0)
        return m('section', {class: 'analysis-section'}, [
            m('h2', r.Name),
            m('div', {id: 'chart-' + id}),
        ])
    },
    onupdate: function(vnode) {
        const id = vnode.attrs.id
        const r = vnode.attrs.data
        const opts = {
            chart: {type: 'area'},
            series: Object.entries(r.Ys).map(function(kv) {
                return {name: kv[0], data: kv[1]}

            }),
            xaxis: {categories: r.X, type: 'datetime'},
        }
        if (this.chart) this.chart.destroy()
        this.chart = new ApexCharts(document.querySelector('#chart-' + id), opts)
        this.chart.render()
    },
}


const App = {
    view: function() {
        if (!$state.data) return
        return [
            m('h1', 'Log Analysis'),
            $state.data.map(
                (r, i) => m(AnalysisSection, {data: r, id: i})
            ),
        ]
    }
}


async function loaddata() {
    let r = await fetch('data.json')
    r = await r.json()
    $state.data = r
    m.redraw()
}


window.onload = function() {
    loaddata()
    m.mount(document.querySelector('main'), App)
}

window.$state = $state
window.m = m
