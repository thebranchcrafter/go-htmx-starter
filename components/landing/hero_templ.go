// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package landing

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Hero() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!-- Hero Section --><section class=\"relative h-screen flex items-center justify-center overflow-hidden bg-gray-900 text-white\"><!-- Parallax Layers --><div class=\"absolute rellax z-10 top-5 left-5 text-blue-600 text-6xl font-bold opacity-10\" data-rellax-speed=\"-8\">Professional</div><div class=\"absolute rellax z-10 top-40 left-1/3 text-yellow-500 text-4xl font-medium opacity-20\" data-rellax-speed=\"4\">Parallax Design</div><div class=\"absolute rellax z-0 w-96 h-96 bg-gradient-to-r from-green-400 to-blue-600 rounded-full opacity-20 top-10 left-10\" data-rellax-speed=\"-6\"></div><div class=\"absolute rellax z-0 w-64 h-64 bg-gradient-to-tl from-pink-500 to-yellow-500 rounded-full opacity-30 bottom-20 right-20\" data-rellax-speed=\"3\"></div><div class=\"absolute rellax z-0 w-48 h-48 bg-gradient-to-br from-purple-400 to-indigo-600 rounded-full opacity-40 top-20 right-1/4\" data-rellax-speed=\"5\"></div><!-- Main Content --><div class=\"relative z-20 text-center px-6\"><h1 class=\"text-6xl md:text-8xl font-extrabold leading-tight tracking-wide\">Elevate Your <span class=\"text-blue-400\">Experience</span></h1><p class=\"mt-6 text-xl text-gray-300\">Discover the perfect blend of innovation, creativity, and functionality in web design.</p><div class=\"mt-10 flex justify-center space-x-6\"><a href=\"#features\" class=\"px-8 py-4 bg-blue-500 text-white rounded-lg shadow-lg hover:bg-blue-600 transition\">Get Started</a> <a href=\"#services\" class=\"px-8 py-4 bg-gray-800 text-white rounded-lg hover:bg-gray-700 transition\">Explore Features</a></div></div><!-- Decorative Elements --><div class=\"absolute rellax z-0 inset-0 bg-gradient-to-br from-blue-600 via-purple-700 to-gray-900 opacity-20\" data-rellax-speed=\"-10\"></div><div class=\"absolute rellax z-0 w-72 h-72 border border-white opacity-10 rounded-full top-10 left-1/2 transform -translate-x-1/2\" data-rellax-speed=\"2\"></div><div class=\"absolute rellax z-0 w-80 h-80 border-2 border-dashed border-gray-400 opacity-15 rounded-full bottom-1/4 right-1/4\" data-rellax-speed=\"-5\"></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
