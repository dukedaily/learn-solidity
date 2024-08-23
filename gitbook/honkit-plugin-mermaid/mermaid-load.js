require([
    'gitbook',
    'jquery'
], function (gitbook, $) {
    $(document).ready(function () {
        mermaid.initialize({startOnLoad: false})
    })
    gitbook.events.bind('page.change', function () {
        $(".mermaid").each(function (i, e) {
            const div = $(e);
            if (div.find('svg').length > 0) return;
            const html = div.html();
            console.log(html)
            const def = html.replace(/&gt;/g, '>').replace(/&lt;/g, '<')
            div.html(render(def));
        });
    })

    const render = function (def) {
        const id = 'mermaid_' + randomString(20);
        const element = document.createElement('dev');
        element.id = id
        element.hidden = true
        return mermaid.mermaidAPI.render(id, def);
    }

    const randomString = function (length) {
        var str = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';
        var result = '';
        for (var i = length; i > 0; --i)
            result += str[Math.floor(Math.random() * str.length)];
        return result;
    }
});
