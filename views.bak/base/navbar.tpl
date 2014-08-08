{{define "navbar"}}
                    <div class="navbar-header">
                        <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
                            <span class="sr-only">Toggle navigation</span>
                            <span class="icon-bar"></span>
                            <span class="icon-bar"></span>
                            <span class="icon-bar"></span>
                        </button>
                        <a class="navbar-brand" href="/">
                            <img src="/static/img/doolp_logo_01.png" alt="Doolp Logo, Changing the world!" width="128" height="52">
                        </a>
                     </div>
                     <div class="navbar-collapse collapse">
                        <ul class="nav navbar-nav navbar-right">

                            <li {{if .IsFeatures}}class="active"{{end}}>
                                <a href="/features/">Features</a>
                            </li>

                            <li {{if .IsPricing}}class="active"{{end}}>
                                <a href="/pricing/">Pricing</a>
                            </li>

                            <li {{if .IsDocmentation}}class="active"{{end}}>
                                <a href="/docs">Documentation</a>
                            </li>

                            <li {{if .IsAPI}}class="active"{{end}}>
                                <a href="/learn">API</a>
                            </li>

                            <li {{if .IsBlog}}class="active"{{end}}>
                                <a href="/blog">Blog</a>
                            </li>

                            <li {{if .IsConsole}}class="active"{{end}}>
                                <a href="/console/">Console</a>
                            </li>
                            
                            {{if .IsLogin}}
                                <li>
                                    <a href="/login?exit=true"><strong>Log out</strong></a>
                                </li>
                            {{else}}
                                <li {{if .IsLogin}}class="active"{{end}}>
                                    <a href="/login"><strong>Log in</strong></a>
                                </li>
                            {{end}}
                        </ul>
                    </div>

                </div>
{{end}}