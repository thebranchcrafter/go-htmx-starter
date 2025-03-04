// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import user_domain "github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user/domain"

func DashboardLayout(u *user_domain.User) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Dashboard</title><link href=\"https://cdn.jsdelivr.net/npm/daisyui@2.6.0/dist/full.css\" rel=\"stylesheet\" type=\"text/css\"><link rel=\"stylesheet\" href=\"/assets/dist/tailwind.min.css\"><script src=\"https://cdn.jsdelivr.net/npm/alpinejs@3.10.2/dist/cdn.min.js\" defer></script><script src=\"https://unpkg.com/htmx.org@1.6.1\"></script><script src=\"https://cdn.jsdelivr.net/npm/chart.js\"></script></head><body class=\"bg-gray-100 font-sans antialiased flex flex-col overscroll-behavior-contain\"><div x-data=\"{ open: true }\" class=\"flex h-max\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Sidebar().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<!-- Main Content --><div class=\"flex-1 flex flex-col bg-gray-100 overflow-y-auto\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Header().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<!-- Page Content --><div class=\"flex-1 p-6 bg-gray-50\"><div class=\"text-gray-700\"><h2 class=\"text-lg font-semibold\">Welcome, ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(u.Name())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/dashboard/dashboard-layout.templ`, Line: 31, Col: 60}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "!</h2><p>Here is your dashboard overview.</p></div><!-- Graph Section --><div class=\"grid grid-cols-1 md:grid-cols-2 gap-6 mt-6\"><!-- Line Chart --><div class=\"bg-white p-4 rounded-lg shadow-md\"><h3 class=\"text-lg font-semibold mb-4\">Monthly Sales</h3><canvas id=\"lineChart\"></canvas></div><!-- Bar Chart --><div class=\"bg-white p-4 rounded-lg shadow-md\"><h3 class=\"text-lg font-semibold mb-4\">Product Performance</h3><canvas id=\"barChart\"></canvas></div><!-- Pie Chart --><div class=\"bg-white p-4 rounded-lg shadow-md\"><h3 class=\"text-lg font-semibold mb-4\">Market Share</h3><canvas id=\"pieChart\"></canvas></div></div></div></div></div><script>\n\t\t\t\t// Line Chart\n\t\t\t\tconst lineCtx = document.getElementById('lineChart').getContext('2d');\n\t\t\t\tnew Chart(lineCtx, {\n\t\t\t\t\ttype: 'line',\n\t\t\t\t\tdata: {\n\t\t\t\t\t\tlabels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul'],\n\t\t\t\t\t\tdatasets: [{\n\t\t\t\t\t\t\tlabel: 'Sales',\n\t\t\t\t\t\t\tdata: [30, 45, 28, 50, 60, 55, 70],\n\t\t\t\t\t\t\tborderColor: 'rgba(75, 192, 192, 1)',\n\t\t\t\t\t\t\tbackgroundColor: 'rgba(75, 192, 192, 0.2)',\n\t\t\t\t\t\t\tfill: true\n\t\t\t\t\t\t}]\n\t\t\t\t\t},\n\t\t\t\t\toptions: {\n\t\t\t\t\t\tresponsive: true,\n\t\t\t\t\t\tplugins: {\n\t\t\t\t\t\t\tlegend: { position: 'top' },\n\t\t\t\t\t\t}\n\t\t\t\t\t}\n\t\t\t\t});\n\n\t\t\t\t// Bar Chart\n\t\t\t\tconst barCtx = document.getElementById('barChart').getContext('2d');\n\t\t\t\tnew Chart(barCtx, {\n\t\t\t\t\ttype: 'bar',\n\t\t\t\t\tdata: {\n\t\t\t\t\t\tlabels: ['Product A', 'Product B', 'Product C', 'Product D'],\n\t\t\t\t\t\tdatasets: [{\n\t\t\t\t\t\t\tlabel: 'Sales',\n\t\t\t\t\t\t\tdata: [20, 35, 50, 25],\n\t\t\t\t\t\t\tbackgroundColor: [\n\t\t\t\t\t\t\t\t'rgba(255, 99, 132, 0.5)',\n\t\t\t\t\t\t\t\t'rgba(54, 162, 235, 0.5)',\n\t\t\t\t\t\t\t\t'rgba(255, 206, 86, 0.5)',\n\t\t\t\t\t\t\t\t'rgba(75, 192, 192, 0.5)'\n\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\tborderColor: [\n\t\t\t\t\t\t\t\t'rgba(255, 99, 132, 1)',\n\t\t\t\t\t\t\t\t'rgba(54, 162, 235, 1)',\n\t\t\t\t\t\t\t\t'rgba(255, 206, 86, 1)',\n\t\t\t\t\t\t\t\t'rgba(75, 192, 192, 1)'\n\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\tborderWidth: 1\n\t\t\t\t\t\t}]\n\t\t\t\t\t},\n\t\t\t\t\toptions: {\n\t\t\t\t\t\tresponsive: true,\n\t\t\t\t\t\tplugins: {\n\t\t\t\t\t\t\tlegend: { position: 'top' },\n\t\t\t\t\t\t}\n\t\t\t\t\t}\n\t\t\t\t});\n\n\t\t\t\t// Pie Chart\n\t\t\t\tconst pieCtx = document.getElementById('pieChart').getContext('2d');\n\t\t\t\tnew Chart(pieCtx, {\n\t\t\t\t\ttype: 'pie',\n\t\t\t\t\tdata: {\n\t\t\t\t\t\tlabels: ['Region A', 'Region B', 'Region C'],\n\t\t\t\t\t\tdatasets: [{\n\t\t\t\t\t\t\tlabel: 'Market Share',\n\t\t\t\t\t\t\tdata: [40, 35, 25],\n\t\t\t\t\t\t\tbackgroundColor: [\n\t\t\t\t\t\t\t\t'rgba(255, 99, 132, 0.5)',\n\t\t\t\t\t\t\t\t'rgba(54, 162, 235, 0.5)',\n\t\t\t\t\t\t\t\t'rgba(255, 206, 86, 0.5)'\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t}]\n\t\t\t\t\t},\n\t\t\t\t\toptions: {\n\t\t\t\t\t\tresponsive: true,\n\t\t\t\t\t\tplugins: {\n\t\t\t\t\t\t\tlegend: { position: 'top' },\n\t\t\t\t\t\t}\n\t\t\t\t\t}\n\t\t\t\t});\n\t\t\t</script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
