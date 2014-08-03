{{define "navbar_console"}}
                    <div class="navbar-header">
                        <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
                            <span class="sr-only">Toggle navigation</span>
                            <span class="icon-bar"></span>
                            <span class="icon-bar"></span>
                            <span class="icon-bar"></span>
                        </button>
                        <a class="navbar-brand" href="/">
                            <img src="/static/img/doolp-logo.png" alt="Doolp Logo, Changing the world!">
                        </a>
                     </div>
                     <div class="navbar-collapse collapse">
                        <ul class="nav navbar-nav navbar-right">

                            <li {{if .IsPricing}}class="active"{{end}}>
                                <a href="/pricing/">Pricing</a>
                            </li>

                            <li {{if .IsDocmentation}}class="active"{{end}}>
                                <a href="/docs">Documentation</a>
                            </li>

                            <li {{if .IsAPI}}class="active"{{end}}>
                                <a href="/learn">API</a>
                            </li>

                            <li>
                                <a href="/login?exit=true"><strong>Log out</strong></a>
                            </li>

                        </ul>
                    </div>

                </div>
{{end}}