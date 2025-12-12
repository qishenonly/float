// Load the complete HTML content directly
function loadPage() {
    const htmlContent = `
<style>
        .text-glow {
            text-shadow: 0 0 30px rgba(79, 70, 229, 0.2);
        }
        /* 稍微减慢流动速度，配合更小的字号显得更沉稳 */
        .animate-gradient-text {
            background-size: 200% auto;
            animation: textShine 6s linear infinite;
        }
        @keyframes textShine {
            to { background-position: 200% center; }
        }
    </style>

    <div class="mesh-bg">
        <div class="blob blob-1"></div>
        <div class="blob blob-2"></div>
        <div class="blob blob-3"></div>
    </div>

    <nav class="fixed top-0 w-full z-50 transition-all duration-300 py-4" id="navbar">
        <div class="container mx-auto px-6">
            <div class="glass rounded-full px-6 py-3 flex justify-between items-center">
                <a href="#" class="flex items-center gap-2 group">
                    <div class="w-8 h-8 rounded-xl bg-gradient-to-tr from-blue-500 to-indigo-600 flex items-center justify-center text-white shadow-lg group-hover:rotate-12 transition-transform">
                        <i class="fa-solid fa-mountain-sun text-sm"></i>
                    </div>
                    <span class="text-lg font-bold text-gray-800 tracking-tight">Float<span class="text-indigo-600">.Island</span></span>
                </a>

                <div class="hidden md:flex items-center gap-8 text-sm font-medium text-gray-500">
                    <a href="#features" class="hover:text-indigo-600 transition">特性</a>
                    <a href="#showcase" class="hover:text-indigo-600 transition">预览</a>
                    <a href="#reviews" class="hover:text-indigo-600 transition">口碑</a>
                    <a href="#faq" class="hover:text-indigo-600 transition">常见问题</a>
                </div>

                <button id="download-btn" class="bg-gray-900 text-white px-5 py-2 rounded-full text-sm font-bold hover:bg-indigo-600 hover:shadow-lg hover:shadow-indigo-200 transition-all transform hover:-translate-y-0.5">
                    下载 App
                </button>
            </div>
        </div>
    </nav>

    <header class="min-h-[110vh] flex items-center justify-center relative pt-20 overflow-hidden">
        <div class="container mx-auto px-6 grid lg:grid-cols-2 gap-16 items-center relative z-10 perspective-wrap">

            <div class="space-y-10 pl-4 lg:pl-0 relative">
                
                <div class="absolute -top-24 -left-10 text-[8rem] font-black text-indigo-50 opacity-40 select-none -z-10 leading-none pointer-events-none">
                    FLOAT
                </div>

                <div class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-white/80 border border-white/50 backdrop-blur-md shadow-sm animate-enter">
                    <span class="flex h-2 w-2 relative">
                        <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-indigo-400 opacity-75"></span>
                        <span class="relative inline-flex rounded-full h-2 w-2 bg-indigo-500"></span>
                    </span>
                    <span class="text-xs font-bold text-indigo-900 uppercase tracking-wide">全新 1.0 版本现已发布</span>
                </div>

                <h1 class="text-4xl lg:text-6xl font-black leading-tight text-gray-900 tracking-tight relative">
                    <span class="block">
                        把生活过得
                        <span class="relative inline-block">
                            <span class="bg-clip-text text-transparent bg-gradient-to-r from-blue-500 via-indigo-500 to-purple-500 animate-gradient-text text-glow">清澈</span>
                            <svg class="absolute w-full h-2 -bottom-1 left-0 text-blue-300 opacity-50" viewBox="0 0 100 10" preserveAspectRatio="none">
                                <path d="M0 5 Q 50 10 100 5" stroke="currentColor" stroke-width="3" fill="none" />
                            </svg>
                        </span>，
                    </span>

                    <span class="block mt-3 ml-12 lg:ml-24">
                        把财富看得
                        <span class="relative inline-block">
                            <span class="bg-clip-text text-transparent bg-gradient-to-r from-purple-500 via-pink-500 to-red-500 animate-gradient-text text-glow">透明</span>。
                            <span class="absolute -top-1 -right-3 w-2.5 h-2.5 bg-pink-400 rounded-full animate-bounce"></span>
                        </span>
                    </span>
                </h1>

                <p class="text-lg text-gray-500 max-w-lg leading-relaxed border-l-4 border-indigo-100 pl-6">
                    浮岛 (Float) 摒弃了枯燥的表格，用<strong>玻璃拟态美学</strong>和<strong>流畅的微交互</strong>，为你构建一座清澈透明的财富岛屿。
                </p>

                <div class="flex flex-wrap gap-4 pt-4">
                    <a href="#" class="download-link flex items-center gap-3 bg-gray-900 text-white px-8 py-4 rounded-2xl hover:scale-105 transition-transform shadow-xl shadow-gray-200 group">
                        <i class="fa-brands fa-apple text-3xl group-hover:text-gray-300 transition"></i>
                        <div class="text-left">
                            <p class="text-[10px] text-gray-400 font-medium tracking-wide uppercase">Download on the</p>
                            <p class="text-base font-bold">App Store</p>
                        </div>
                    </a>
                    <a href="#" class="download-link flex items-center gap-3 bg-white text-gray-900 border border-gray-200 px-8 py-4 rounded-2xl hover:bg-gray-50 transition hover:scale-105 group">
                        <i class="fa-brands fa-android text-3xl text-green-500 group-hover:scale-110 transition"></i>
                        <div class="text-left">
                            <p class="text-[10px] text-gray-400 font-medium tracking-wide uppercase">Get it on</p>
                            <p class="text-base font-bold">Google Play</p>
                        </div>
                    </a>
                </div>

                <div class="flex items-center gap-4 pt-2">
                    <div class="flex -space-x-3">
                        <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=Felix" class="w-10 h-10 rounded-full border-[3px] border-white bg-gray-100 shadow-sm" alt="">
                        <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=Aneka" class="w-10 h-10 rounded-full border-[3px] border-white bg-gray-100 shadow-sm" alt="">
                        <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=Zack" class="w-10 h-10 rounded-full border-[3px] border-white bg-gray-100 shadow-sm" alt="">
                    </div>
                    <div>
                        <div class="flex items-center gap-1 text-yellow-500 text-xs mb-0.5">
                            <i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i>
                        </div>
                        <p class="text-sm font-bold text-gray-600">50,000+ 岛民已入住</p>
                    </div>
                </div>
            </div>

            <div class="h-[800px] flex items-center justify-center relative" id="hero-tilt-container">
                <div class="tilt-card relative w-[360px] h-[720px]" id="hero-phone">
                    <div class="absolute top-[15%] -left-[15%] z-50 transform translate-z-20 float-anim">
                        <div class="glass-strong p-4 rounded-2xl flex items-center gap-4 shadow-float w-64 border-l-4 border-indigo-500">
                            <div class="w-12 h-12 rounded-xl bg-indigo-50 text-indigo-600 flex items-center justify-center text-xl"><i class="fa-solid fa-wallet"></i></div>
                            <div>
                                <p class="text-xs text-gray-400 font-bold uppercase">Net Worth</p>
                                <p class="text-xl font-bold text-gray-800">¥ 124,500</p>
                            </div>
                        </div>
                    </div>

                    <div class="absolute bottom-[20%] -right-[10%] z-50 transform translate-z-30 float-anim" style="animation-delay: 1.5s;">
                        <div class="glass-strong p-4 rounded-2xl flex items-center gap-3 shadow-float w-56 border-l-4 border-red-400">
                            <div class="w-10 h-10 rounded-full bg-black text-red-500 flex items-center justify-center text-lg"><i class="fa-brands fa-netflix"></i></div>
                            <div>
                                <p class="text-sm font-bold text-gray-800">Netflix</p>
                                <p class="text-xs text-red-500 font-bold">明日扣费 ¥39</p>
                            </div>
                        </div>
                    </div>

                    <div class="w-full h-full bg-white rounded-[50px] border-[8px] border-gray-900 shadow-2xl overflow-hidden relative z-20">
                        <div class="absolute top-0 left-1/2 transform -translate-x-1/2 w-32 h-7 bg-gray-900 rounded-b-2xl z-50"></div>

                        <div class="w-full h-full bg-gradient-to-br from-blue-50 via-white to-purple-50 flex flex-col pt-16 px-6 relative">
                            <div class="absolute top-20 right-[-20px] w-32 h-32 bg-purple-200 rounded-full mix-blend-multiply filter blur-xl opacity-70"></div>
                            <div class="absolute top-40 left-[-20px] w-32 h-32 bg-blue-200 rounded-full mix-blend-multiply filter blur-xl opacity-70"></div>

                            <div class="relative z-10">
                                <h2 class="text-2xl font-extrabold text-gray-800 mb-1">Hi, Alex 👋</h2>
                                <p class="text-sm text-gray-400 mb-8">今天也是充满希望的一天 ✨</p>

                                <div class="bg-gradient-to-tr from-indigo-600 to-purple-600 rounded-3xl p-6 text-white shadow-xl shadow-indigo-200 mb-6">
                                    <p class="text-indigo-100 text-xs font-medium mb-1">本月结余</p>
                                    <h3 class="text-3xl font-bold mb-4">¥ 4,250.00</h3>
                                    <div class="w-full bg-black/20 h-1.5 rounded-full overflow-hidden">
                                        <div class="bg-white h-full w-[45%]"></div>
                                    </div>
                                </div>

                                <div class="space-y-4">
                                    <div class="bg-white/60 backdrop-blur-md p-4 rounded-2xl flex justify-between items-center shadow-sm">
                                        <div class="flex items-center gap-3">
                                            <div class="w-10 h-10 rounded-xl bg-orange-50 text-orange-500 flex items-center justify-center"><i class="fa-solid fa-burger"></i></div>
                                            <div><p class="font-bold text-gray-800 text-sm">Shake Shack</p><p class="text-xs text-gray-400">晚餐</p></div>
                                        </div>
                                        <span class="font-bold text-gray-800">- 85.00</span>
                                    </div>
                                    <div class="bg-white/60 backdrop-blur-md p-4 rounded-2xl flex justify-between items-center shadow-sm">
                                        <div class="flex items-center gap-3">
                                            <div class="w-10 h-10 rounded-xl bg-blue-50 text-blue-500 flex items-center justify-center"><i class="fa-solid fa-train-subway"></i></div>
                                            <div><p class="font-bold text-gray-800 text-sm">地铁通勤</p><p class="text-xs text-gray-400">交通</p></div>
                                        </div>
                                        <span class="font-bold text-gray-800">- 5.00</span>
                                    </div>
                                </div>
                            </div>

                            <div class="absolute bottom-0 left-0 w-full bg-white/80 backdrop-blur-xl border-t border-gray-100 p-4 pb-8 flex justify-between px-8 text-gray-400 text-xl">
                                <i class="fa-solid fa-house text-gray-800"></i>
                                <i class="fa-solid fa-chart-pie"></i>
                                <div class="w-12 h-12 bg-gray-900 rounded-full text-white flex items-center justify-center -mt-6 border-4 border-white shadow-lg"><i class="fa-solid fa-plus"></i></div>
                                <i class="fa-solid fa-wallet"></i>
                                <i class="fa-solid fa-user"></i>
                            </div>
                        </div>
                    </div>

                    <div class="absolute -bottom-10 left-10 right-10 h-10 bg-black/20 blur-2xl rounded-[100%] z-0"></div>
                </div>
            </div>
        </div>
    </header>
    
    <section id="features" class="py-24 bg-white/50 relative z-10">
        <div class="container mx-auto px-6">

            <div class="flex flex-col md:flex-row items-center gap-16 mb-32 feature-block">
                <div class="flex-1 order-2 md:order-1 relative">
                    <div class="absolute inset-0 bg-blue-100 rounded-full filter blur-3xl opacity-50 transform -translate-x-10"></div>
                    <div class="glass-strong rounded-[2.5rem] p-8 relative shadow-2xl transform rotate-[-2deg] hover:rotate-0 transition duration-500">
                        <div class="flex justify-between items-center mb-6">
                            <h3 class="font-bold text-gray-800">资金账户</h3>
                            <div class="w-8 h-8 bg-gray-100 rounded-full flex items-center justify-center text-gray-400"><i class="fa-solid fa-plus"></i></div>
                        </div>
                        <div class="space-y-4">
                            <div class="bg-white p-4 rounded-2xl flex items-center justify-between shadow-sm border border-gray-50">
                                <div class="flex gap-4">
                                    <div class="w-10 h-10 rounded-xl bg-red-50 text-red-500 flex items-center justify-center"><i class="fa-solid fa-building-columns"></i></div>
                                    <div><p class="font-bold text-sm">招商银行</p><p class="text-xs text-gray-400">储蓄卡 • 8899</p></div>
                                </div>
                                <span class="font-bold">¥ 85,200</span>
                            </div>
                            <div class="bg-white p-4 rounded-2xl flex items-center justify-between shadow-sm border border-gray-50">
                                <div class="flex gap-4">
                                    <div class="w-10 h-10 rounded-xl bg-blue-50 text-blue-500 flex items-center justify-center"><i class="fa-brands fa-alipay text-lg"></i></div>
                                    <div><p class="font-bold text-sm">支付宝</p><p class="text-xs text-gray-400">余额宝</p></div>
                                </div>
                                <span class="font-bold">¥ 42,300</span>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="flex-1 order-1 md:order-2 space-y-6">
                    <div class="w-12 h-12 rounded-2xl bg-blue-100 text-blue-600 flex items-center justify-center text-2xl mb-4"><i class="fa-solid fa-wallet"></i></div>
                    <h2 class="text-4xl font-bold text-gray-900">资产全景，<br>尽在掌握。</h2>
                    <p class="text-gray-500 text-lg leading-relaxed">
                        告别在各个银行 App 之间切换的烦恼。浮岛通过直观的卡片设计，将您的储蓄、现金、负债整合在一处。看着数字增长，本身就是一种享受。
                    </p>
                    <ul class="space-y-3 text-gray-600 font-medium">
                        <li class="flex items-center gap-3"><i class="fa-solid fa-check text-green-500"></i> 自动计算净资产</li>
                        <li class="flex items-center gap-3"><i class="fa-solid fa-check text-green-500"></i> 支持信用卡/花呗负债管理</li>
                    </ul>
                </div>
            </div>

            <div class="flex flex-col md:flex-row items-center gap-16 mb-32 feature-block">
                <div class="flex-1 space-y-6">
                    <div class="w-12 h-12 rounded-2xl bg-orange-100 text-orange-500 flex items-center justify-center text-2xl mb-4"><i class="fa-solid fa-bolt"></i></div>
                    <h2 class="text-4xl font-bold text-gray-900">3 秒极速记账，<br>不打断生活。</h2>
                    <p class="text-gray-500 text-lg leading-relaxed">
                        我们优化了记账流程的每一个毫秒。无启动广告，无繁琐步骤。打开即记，选类即存。就像发一条微信一样简单。
                    </p>
                    <ul class="space-y-3 text-gray-600 font-medium">
                        <li class="flex items-center gap-3"><i class="fa-solid fa-check text-green-500"></i> 智能分类预测</li>
                        <li class="flex items-center gap-3"><i class="fa-solid fa-check text-green-500"></i> 自定义常用类别</li>
                    </ul>
                </div>
                <div class="flex-1 relative">
                     <div class="absolute inset-0 bg-orange-100 rounded-full filter blur-3xl opacity-50 transform translate-x-10"></div>
                     <div class="glass-strong rounded-[2.5rem] p-8 relative shadow-2xl transform rotate-[2deg] hover:rotate-0 transition duration-500">
                         <div class="flex justify-end mb-8">
                             <div class="text-right">
                                 <div class="flex items-center justify-end gap-2 text-gray-500 mb-1">
                                     <div class="w-6 h-6 rounded-full bg-orange-100 text-orange-500 flex items-center justify-center text-xs"><i class="fa-solid fa-utensils"></i></div>
                                     <span class="text-sm">餐饮美食</span>
                                 </div>
                                 <div class="text-4xl font-bold text-gray-800">48.00</div>
                             </div>
                         </div>
                         <div class="grid grid-cols-4 gap-3">
                             <div class="aspect-square bg-gray-50 rounded-xl flex items-center justify-center text-xl font-bold text-gray-700 hover:bg-gray-200 transition">1</div>
                             <div class="aspect-square bg-gray-50 rounded-xl flex items-center justify-center text-xl font-bold text-gray-700 hover:bg-gray-200 transition">2</div>
                             <div class="aspect-square bg-gray-50 rounded-xl flex items-center justify-center text-xl font-bold text-gray-700 hover:bg-gray-200 transition">3</div>
                             <div class="bg-gray-100 rounded-xl flex items-center justify-center text-gray-500 hover:bg-red-50 hover:text-red-500 transition"><i class="fa-solid fa-delete-left"></i></div>

                             <div class="aspect-square bg-gray-50 rounded-xl flex items-center justify-center text-xl font-bold text-gray-700 hover:bg-gray-200 transition">4</div>
                             <div class="aspect-square bg-gray-50 rounded-xl flex items-center justify-center text-xl font-bold text-gray-700 hover:bg-gray-200 transition">5</div>
                             <div class="aspect-square bg-gray-50 rounded-xl flex items-center justify-center text-xl font-bold text-gray-700 hover:bg-gray-200 transition">6</div>
                             <div class="row-span-2 bg-gray-900 rounded-xl flex items-center justify-center text-white font-bold text-sm shadow-lg hover:bg-gray-800 transition">完成</div>

                             <div class="aspect-square bg-gray-50 rounded-xl flex items-center justify-center text-xl font-bold text-gray-700 hover:bg-gray-200 transition">7</div>
                             <div class="aspect-square bg-gray-50 rounded-xl flex items-center justify-center text-xl font-bold text-gray-700 hover:bg-gray-200 transition">8</div>
                             <div class="aspect-square bg-gray-50 rounded-xl flex items-center justify-center text-xl font-bold text-gray-700 hover:bg-gray-200 transition">9</div>
                         </div>
                     </div>
                </div>
            </div>

        </div>
    </section>

    <section class="py-20 relative">
        <div class="container mx-auto px-6">
            <div class="text-center mb-16">
                <span class="text-indigo-600 font-bold tracking-wider uppercase text-xs mb-2 block">Why Float Island?</span>
                <h2 class="text-4xl font-bold text-gray-900">不仅好看，<br>更是强大的财务中台</h2>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-3 md:grid-rows-2 gap-6 h-auto md:h-[600px]">

                <div class="md:col-span-2 md:row-span-2 glass-strong rounded-[2rem] p-8 group hover:shadow-glass transition duration-300 relative overflow-hidden">
                    <div class="absolute top-0 right-0 w-64 h-64 bg-purple-100 rounded-full filter blur-3xl opacity-60"></div>
                    <div class="relative z-10">
                        <div class="w-12 h-12 rounded-2xl bg-purple-100 text-purple-600 flex items-center justify-center text-2xl mb-6"><i class="fa-solid fa-bell"></i></div>
                        <h3 class="text-2xl font-bold mb-2">订阅雷达</h3>
                        <p class="text-gray-500 mb-8 max-w-sm">别再为忘记取消的会员付费。Float 会在扣费前 24 小时贴心提醒，让你的每一分钱都花在刀刃上。</p>

                        <div class="space-y-3">
                            <div class="bg-white p-3 rounded-xl flex items-center justify-between border-l-4 border-red-500 shadow-sm transform group-hover:translate-x-2 transition">
                                <div class="flex items-center gap-3">
                                    <div class="w-8 h-8 rounded bg-black text-red-600 flex items-center justify-center text-xs"><i class="fa-brands fa-netflix"></i></div>
                                    <div class="text-sm font-bold text-gray-800">Netflix</div>
                                </div>
                                <div class="text-xs text-red-500 font-bold bg-red-50 px-2 py-1 rounded">明天到期</div>
                            </div>
                             <div class="bg-white p-3 rounded-xl flex items-center justify-between border-l-4 border-green-500 shadow-sm transform group-hover:translate-x-2 transition delay-75">
                                <div class="flex items-center gap-3">
                                    <div class="w-8 h-8 rounded bg-green-100 text-green-600 flex items-center justify-center text-xs"><i class="fa-brands fa-spotify"></i></div>
                                    <div class="text-sm font-bold text-gray-800">Spotify</div>
                                </div>
                                <div class="text-xs text-gray-400">3天后</div>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="glass-strong rounded-[2rem] p-6 flex flex-col justify-between group hover:bg-white transition">
                    <div class="w-10 h-10 rounded-xl bg-green-100 text-green-600 flex items-center justify-center text-lg"><i class="fa-solid fa-file-excel"></i></div>
                    <div>
                        <h3 class="font-bold text-lg mb-1">自由导出</h3>
                        <p class="text-xs text-gray-400">一键生成 Excel / PDF 报表，数据掌握在自己手中。</p>
                    </div>
                    <div class="absolute right-4 top-4 opacity-0 group-hover:opacity-100 transition transform group-hover:rotate-45 text-green-500"><i class="fa-solid fa-arrow-up-right-from-square"></i></div>
                </div>

                <div class="glass-strong rounded-[2rem] p-6 flex flex-col justify-between group hover:bg-white transition">
                    <div class="w-10 h-10 rounded-xl bg-blue-100 text-blue-600 flex items-center justify-center text-lg"><i class="fa-solid fa-cloud"></i></div>
                    <div>
                        <h3 class="font-bold text-lg mb-1">云端同步</h3>
                        <p class="text-xs text-gray-400">支持 Google Drive、OneDrive 等云服务，数据实时更新。</p>
                    </div>
                </div>

                <div class="glass-strong rounded-[2rem] p-6 flex items-center gap-4 md:col-span-1 group hover:bg-white transition">
                     <div class="w-14 h-14 rounded-full border-2 border-gray-200 flex items-center justify-center text-2xl text-gray-400 group-hover:text-indigo-600 group-hover:border-indigo-600 transition"><i class="fa-solid fa-fingerprint"></i></div>
                     <div>
                         <h3 class="font-bold text-lg">生物识别隐私锁</h3>
                         <p class="text-xs text-gray-400 mt-1">支持指纹、面部识别，本地加密存储，你的数据只属于你。</p>
                     </div>
                </div>

            </div>
        </div>
    </section>

    <section id="reviews" class="py-20 bg-white border-t border-gray-100">
        <div class="container mx-auto px-6">
            <h2 class="text-3xl font-bold text-center mb-16">听听岛民们怎么说</h2>
            <div class="grid md:grid-cols-3 gap-8">
                <div class="bg-gray-50 p-8 rounded-3xl hover:bg-white hover:shadow-xl transition duration-300">
                    <div class="flex items-center gap-1 text-yellow-400 mb-4 text-xs">
                        <i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i>
                    </div>
                    <p class="text-gray-600 italic mb-6">"终于找到了一款不臃肿的记账软件。界面太美了，每次打开都是一种享受，让我重新爱上了记账。"</p>
                    <div class="flex items-center gap-3">
                        <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=Lily" class="w-10 h-10 rounded-full" alt="">
                        <div><p class="font-bold text-sm">Lily Chen</p><p class="text-xs text-gray-400">UI 设计师</p></div>
                    </div>
                </div>
                 <div class="bg-gray-50 p-8 rounded-3xl hover:bg-white hover:shadow-xl transition duration-300 transform md:-translate-y-4">
                    <div class="flex items-center gap-1 text-yellow-400 mb-4 text-xs">
                        <i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i>
                    </div>
                    <p class="text-gray-600 italic mb-6">"订阅管理功能帮我省了不少钱！发现了很多忘记取消的自动续费。资产全景功能也很赞。"</p>
                    <div class="flex items-center gap-3">
                        <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=Mike" class="w-10 h-10 rounded-full" alt="">
                        <div><p class="font-bold text-sm">Mike Wang</p><p class="text-xs text-gray-400">产品经理</p></div>
                    </div>
                </div>
                 <div class="bg-gray-50 p-8 rounded-3xl hover:bg-white hover:shadow-xl transition duration-300">
                    <div class="flex items-center gap-1 text-yellow-400 mb-4 text-xs">
                        <i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i>
                    </div>
                    <p class="text-gray-600 italic mb-6">"极简主义者的福音。没有广告，没有理财推销，只有纯粹的记账。希望开发团队保持初心。"</p>
                    <div class="flex items-center gap-3">
                        <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah" class="w-10 h-10 rounded-full" alt="">
                        <div><p class="font-bold text-sm">Sarah Li</p><p class="text-xs text-gray-400">自由职业者</p></div>
                    </div>
                </div>
            </div>
        </div>
    </section>

    <section id="faq" class="py-20">
        <div class="container mx-auto px-6 max-w-3xl">
            <h2 class="text-3xl font-bold mb-12 text-center">常见问题</h2>

            <div class="space-y-4">
                <div class="glass p-6 rounded-2xl cursor-pointer group">
                    <div class="flex justify-between items-center">
                        <h4 class="font-bold text-gray-800">数据安全吗？会上传服务器吗？</h4>
                        <i class="fa-solid fa-chevron-down text-gray-400 group-hover:text-indigo-600 transition"></i>
                    </div>
                    <p class="text-gray-500 text-sm mt-3 hidden group-hover:block transition-all">浮岛采用本地优先策略。未开启同步时，数据仅保存在您的手机中。开启云端同步后，数据加密存储于您选择的云服务（如 Google Drive、OneDrive），我们无法访问您的任何隐私数据。</p>
                </div>

                <div class="glass p-6 rounded-2xl cursor-pointer group">
                    <div class="flex justify-between items-center">
                        <h4 class="font-bold text-gray-800">支持导入其他软件的账单吗？</h4>
                        <i class="fa-solid fa-chevron-down text-gray-400 group-hover:text-indigo-600 transition"></i>
                    </div>
                    <p class="text-gray-500 text-sm mt-3 hidden group-hover:block transition-all">支持。我们提供了标准 Excel 模板，您可以将其他软件的数据整理后一键导入。</p>
                </div>

                 <div class="glass p-6 rounded-2xl cursor-pointer group">
                    <div class="flex justify-between items-center">
                        <h4 class="font-bold text-gray-800">是免费的吗？</h4>
                        <i class="fa-solid fa-chevron-down text-gray-400 group-hover:text-indigo-600 transition"></i>
                    </div>
                    <p class="text-gray-500 text-sm mt-3 hidden group-hover:block transition-all">基础记账功能永久免费。高级功能（如无限资产账户、订阅雷达、导出功能）需要订阅 Float Pro 会员。</p>
                </div>
            </div>
        </div>
    </section>

    <footer class="bg-gray-900 text-white pt-20 pb-10">
        <div class="container mx-auto px-6">
            <div class="flex flex-col md:flex-row justify-between items-center mb-16">
                <div class="text-center md:text-left mb-8 md:mb-0">
                    <h2 class="text-3xl font-bold mb-2">Ready to Float?</h2>
                    <p class="text-gray-400">即刻下载，构建你的财富岛屿。</p>
                </div>
                <div class="flex gap-4">
                    <button class="w-12 h-12 rounded-full bg-white/10 flex items-center justify-center hover:bg-white hover:text-black transition"><i class="fa-brands fa-weixin"></i></button>
                    <button class="w-12 h-12 rounded-full bg-white/10 flex items-center justify-center hover:bg-white hover:text-black transition"><i class="fa-brands fa-weibo"></i></button>
                    <button class="w-12 h-12 rounded-full bg-white/10 flex items-center justify-center hover:bg-white hover:text-black transition"><i class="fa-solid fa-envelope"></i></button>
                </div>
            </div>

            <div class="border-t border-gray-800 pt-8 flex flex-col md:flex-row justify-between items-center text-xs text-gray-500">
                <p>&copy; 2025 Float Island Inc. All rights reserved.</p>
                <div class="flex gap-6 mt-4 md:mt-0">
                    <a href="#" id="privacy-link" class="hover:text-white transition cursor-pointer">隐私协议</a>
                    <a href="#" id="terms-link" class="hover:text-white transition cursor-pointer">使用条款</a>
                </div>
            </div>
        </div>
    </footer>

    <!-- Privacy Policy Modal -->
    <div id="privacy-modal" class="fixed inset-0 z-50 hidden bg-black/50 backdrop-blur-sm">
        <div class="flex items-center justify-center min-h-screen p-4">
            <div class="bg-white rounded-3xl max-w-4xl w-full max-h-[90vh] overflow-hidden shadow-2xl">
                <div class="flex justify-between items-center p-6 border-b border-gray-100">
                    <h1 class="text-2xl font-bold text-gray-900">隐私协议</h1>
                    <button id="close-privacy-modal" class="w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center hover:bg-gray-200 transition">
                        <i class="fa-solid fa-x text-gray-500"></i>
                    </button>
                </div>
                <div class="p-8 overflow-y-auto max-h-[calc(90vh-120px)]">
                    <div class="text-center mb-8">
                        <p class="text-gray-600">最后更新时间：2025年12月8日</p>
                    </div>
                    <div class="space-y-8">
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">1. 信息收集与使用</h2>
                            <div class="space-y-4 text-gray-700">
                                <p>浮岛（Float Island）非常重视您的隐私保护。我们致力于保护您的个人信息安全。本隐私协议说明了我们如何收集、使用和保护您的信息。</p>
                                <h3 class="text-lg font-semibold text-gray-800">1.1 我们收集的信息</h3>
                                <ul class="list-disc pl-6 space-y-2">
                                    <li><strong>财务数据：</strong>您手动输入的收支记录、账户余额、预算信息等财务数据</li>
                                    <li><strong>设备信息：</strong>设备型号、操作系统版本、应用版本等基本信息</li>
                                    <li><strong>使用数据：</strong>应用使用频率、功能使用情况等统计信息</li>
                                    <li><strong>崩溃报告：</strong>应用崩溃时的错误日志（用于改进应用稳定性）</li>
                                </ul>
                                <h3 class="text-lg font-semibold text-gray-800">1.2 信息使用目的</h3>
                                <ul class="list-disc pl-6 space-y-2">
                                    <li>提供记账和财务管理核心功能</li>
                                    <li>改进应用性能和用户体验</li>
                                    <li>保障应用安全和稳定性</li>
                                    <li>向您提供个性化建议和提醒</li>
                                </ul>
                            </div>
                        </section>
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">2. 数据存储与安全</h2>
                            <div class="space-y-4 text-gray-700">
                                <h3 class="text-lg font-semibold text-gray-800">2.1 本地存储</h3>
                                <p>您的所有财务数据默认存储在您的设备本地。我们使用行业标准的加密技术保护您的数据安全。</p>
                                <h3 class="text-lg font-semibold text-gray-800">2.2 云端同步（可选）</h3>
                                <p>如果您选择开启云端同步功能，我们会将您的数据加密后存储在云服务器上。数据传输和存储均采用SSL/TLS加密。</p>
                                <h3 class="text-lg font-semibold text-gray-800">2.3 数据安全措施</h3>
                                <ul class="list-disc pl-6 space-y-2">
                                    <li>采用AES-256加密算法保护数据</li>
                                    <li>使用生物识别或PIN码进行应用锁定</li>
                                    <li>定期进行安全审计和更新</li>
                                    <li>严格限制内部人员访问权限</li>
                                </ul>
                            </div>
                        </section>
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">3. 信息共享与披露</h2>
                            <div class="space-y-4 text-gray-700">
                                <p>我们承诺不会出售、出租或以其他方式披露您的个人信息，除非以下情况：</p>
                                <ul class="list-disc pl-6 space-y-2">
                                    <li>获得您的明确同意</li>
                                    <li>法律法规要求或政府部门要求</li>
                                    <li>保护我们的合法权益</li>
                                    <li>涉及合并、收购或资产转让</li>
                                </ul>
                                <p class="text-sm text-gray-600 mt-4">我们不会向第三方广告商或营销公司提供您的个人信息。</p>
                            </div>
                        </section>
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">4. Cookie和类似技术</h2>
                            <div class="space-y-4 text-gray-700">
                                <p>我们的网站可能使用Cookie来改善用户体验。这些Cookie仅用于：</p>
                                <ul class="list-disc pl-6 space-y-2">
                                    <li>记住您的偏好设置</li>
                                    <li>分析网站使用情况</li>
                                    <li>提供基本的网站功能</li>
                                </ul>
                                <p>您可以通过浏览器设置拒绝Cookie，但这可能影响网站的部分功能。</p>
                            </div>
                        </section>
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">5. 未成年人保护</h2>
                            <div class="space-y-4 text-gray-700">
                                <p>我们非常重视未成年人的隐私保护。如果您未满18岁，请在监护人的陪同下使用我们的服务。我们不会故意收集未成年人的个人信息。</p>
                            </div>
                        </section>
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">6. 您的权利</h2>
                            <div class="space-y-4 text-gray-700">
                                <p>根据相关法律法规，您享有以下权利：</p>
                                <ul class="list-disc pl-6 space-y-2">
                                    <li>访问我们收集的关于您的信息</li>
                                    <li>更正不准确或不完整的信息</li>
                                    <li>删除您的个人信息</li>
                                    <li>反对或限制我们处理您的信息</li>
                                    <li>数据可移植性</li>
                                </ul>
                                <p>如需行使这些权利，请通过应用内的设置或联系我们。</p>
                            </div>
                        </section>
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">7. 隐私协议更新</h2>
                            <div class="space-y-4 text-gray-700">
                                <p>我们可能会不定期更新本隐私协议。重大变更时，我们会通过应用内通知或网站公告等方式告知您。</p>
                                <p>继续使用我们的服务即表示您同意更新后的隐私协议。</p>
                            </div>
                        </section>
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">8. 联系我们</h2>
                            <div class="space-y-4 text-gray-700">
                                <p>如果您对本隐私协议有任何疑问、意见或建议，请通过以下方式联系我们：</p>
                                <div class="bg-gray-50 p-4 rounded-lg">
                                    <p><strong>邮箱：</strong>privacy@floatisland.app</p>
                                    <p><strong>网站：</strong><a href="https://floatisland.app" class="text-indigo-600 hover:underline">https://floatisland.app</a></p>
                                </div>
                            </div>
                        </section>
                        <div class="border-t pt-6 mt-8">
                            <p class="text-sm text-gray-600 text-center">
                                本隐私协议适用于浮岛（Float Island）应用及相关服务。<br>
                                最后更新：2025年12月8日
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <!-- Terms of Use Modal -->
    <div id="terms-modal" class="fixed inset-0 z-50 hidden bg-black/50 backdrop-blur-sm">
        <div class="flex items-center justify-center min-h-screen p-4">
            <div class="bg-white rounded-3xl max-w-4xl w-full max-h-[90vh] overflow-hidden shadow-2xl">
                <div class="flex justify-between items-center p-6 border-b border-gray-100">
                    <h1 class="text-2xl font-bold text-gray-900">使用条款</h1>
                    <button id="close-terms-modal" class="w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center hover:bg-gray-200 transition">
                        <i class="fa-solid fa-x text-gray-500"></i>
                    </button>
                </div>
                <div class="p-8 overflow-y-auto max-h-[calc(90vh-120px)]">
                    <div class="text-center mb-8">
                        <p class="text-gray-600">最后更新时间：2025年12月8日</p>
                    </div>
                    <div class="space-y-8">
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">1. 接受条款</h2>
                            <div class="space-y-4 text-gray-700">
                                <p>欢迎使用浮岛（Float Island）！本使用条款（以下简称"条款"）是您与Float Island Inc.（以下简称"我们"或"公司"）之间的法律协议。</p>
                                <p>通过下载、安装或使用浮岛应用，您即表示您已阅读、理解并同意受本条款约束。如果您不同意这些条款，请不要使用我们的应用。</p>
                            </div>
                        </section>
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">2. 服务描述</h2>
                            <div class="space-y-4 text-gray-700">
                                <p>浮岛是一款个人财务管理应用，提供以下核心功能：</p>
                                <ul class="list-disc pl-6 space-y-2">
                                    <li>收支记录和分类管理</li>
                                    <li>资产负债全景展示</li>
                                    <li>预算制定和监控</li>
                                    <li>财务报表生成</li>
                                    <li>订阅服务提醒</li>
                                    <li>数据导入导出功能</li>
                                </ul>
                            </div>
                        </section>
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">3. 用户资格</h2>
                            <div class="space-y-4 text-gray-700">
                                <p>要使用浮岛服务，您必须：</p>
                                <ul class="list-disc pl-6 space-y-2">
                                    <li>年满13岁（或您所在司法管辖区的法定成年年龄）</li>
                                    <li>能够理解并同意本条款</li>
                                    <li>遵守所有适用的法律法规</li>
                                </ul>
                                <p>如果您未达到法定年龄，请在监护人的指导下使用本服务。</p>
                            </div>
                        </section>
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">4. 用户责任</h2>
                            <div class="space-y-4 text-gray-700">
                                <h3 class="text-lg font-semibold text-gray-800">4.1 账户安全</h3>
                                <p>您负责维护您账户和密码的安全性。您同意：</p>
                                <ul class="list-disc pl-6 space-y-2">
                                    <li>不与他人分享您的账户信息</li>
                                    <li>及时更新您的个人信息</li>
                                    <li>在发现未经授权使用时立即通知我们</li>
                                </ul>
                                <h3 class="text-lg font-semibold text-gray-800">4.2 数据准确性</h3>
                                <p>您负责确保输入到应用中的财务数据的准确性和完整性。我们不对基于不准确数据做出的财务决策承担责任。</p>
                                <h3 class="text-lg font-semibold text-gray-800">4.3 合规使用</h3>
                                <p>您同意仅将本应用用于合法目的，不得：</p>
                                <ul class="list-disc pl-6 space-y-2">
                                    <li>违反任何适用的法律法规</li>
                                    <li>侵犯他人知识产权</li>
                                    <li>传播恶意软件或病毒</li>
                                    <li>进行任何形式的欺诈活动</li>
                                </ul>
                            </div>
                        </section>
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">5. 付费服务</h2>
                            <div class="space-y-4 text-gray-700">
                                <h3 class="text-lg font-semibold text-gray-800">5.1 免费功能</h3>
                                <p>浮岛提供基础的记账功能永久免费使用。</p>
                                <h3 class="text-lg font-semibold text-gray-800">5.2 高级功能</h3>
                                <p>部分高级功能需要订阅Float Pro会员，包括：</p>
                                <ul class="list-disc pl-6 space-y-2">
                                    <li>无限资产账户管理</li>
                                    <li>订阅服务智能提醒</li>
                                    <li>高级财务报表</li>
                                    <li>数据导出功能</li>
                                    <li>云端同步</li>
                                </ul>
                                <h3 class="text-lg font-semibold text-gray-800">5.3 付费条款</h3>
                                <ul class="list-disc pl-6 space-y-2">
                                    <li>订阅费用以应用内显示为准</li>
                                    <li>订阅自动续费，可随时取消</li>
                                    <li>取消后仍可在付费期间使用高级功能</li>
                                    <li>所有销售均为最终销售，不予退款</li>
                                </ul>
                            </div>
                        </section>
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">6. 知识产权</h2>
                            <div class="space-y-4 text-gray-700">
                                <p>浮岛应用及其所有内容（包括但不限于软件、设计、商标、文本、图像等）的知识产权归Float Island Inc.所有。</p>
                                <p>您被授予有限的、非排他性的、不可转让的使用许可，仅用于个人非商业目的。您不得：</p>
                                <ul class="list-disc pl-6 space-y-2">
                                    <li>复制、分发或修改应用</li>
                                    <li>逆向工程或反编译应用</li>
                                    <li>移除版权声明或标识</li>
                                    <li>将应用用于商业目的</li>
                                </ul>
                            </div>
                        </section>
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">7. 免责声明</h2>
                            <div class="space-y-4 text-gray-700">
                                <p>浮岛服务按"现状"提供，我们不保证：</p>
                                <ul class="list-disc pl-6 space-y-2">
                                    <li>服务不会中断或无错误</li>
                                    <li>所有功能将始终可用</li>
                                    <li>应用完全安全无漏洞</li>
                                    <li>财务建议的准确性</li>
                                </ul>
                                <p class="text-sm text-gray-600 mt-4">我们不对因使用本应用而产生的任何直接、间接、偶然或后果性损害承担责任。</p>
                            </div>
                        </section>
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">8. 服务终止</h2>
                            <div class="space-y-4 text-gray-700">
                                <p>我们保留随时终止或暂停您使用服务的权利，包括但不限于：</p>
                                <ul class="list-disc pl-6 space-y-2">
                                    <li>违反本使用条款</li>
                                    <li>违反适用法律法规</li>
                                    <li>滥用服务功能</li>
                                    <li>技术或安全原因</li>
                                </ul>
                                <p>终止服务后，我们可能会删除您的账户和相关数据。</p>
                            </div>
                        </section>
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">9. 条款修改</h2>
                            <div class="space-y-4 text-gray-700">
                                <p>我们保留随时修改本使用条款的权利。重大变更时，我们会通过应用内通知或网站公告等方式提前告知您。</p>
                                <p>继续使用我们的服务即表示您接受修改后的条款。</p>
                            </div>
                        </section>
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">10. 适用法律</h2>
                            <div class="space-y-4 text-gray-700">
                                <p>本条款受中华人民共和国法律管辖并依据其解释。如发生争议，应提交至有管辖权的人民法院解决。</p>
                            </div>
                        </section>
                        <section>
                            <h2 class="text-xl font-bold text-gray-900 mb-4">11. 联系我们</h2>
                            <div class="space-y-4 text-gray-700">
                                <p>如果您对本使用条款有任何疑问，请通过以下方式联系我们：</p>
                                <div class="bg-gray-50 p-4 rounded-lg">
                                    <p><strong>邮箱：</strong>legal@floatisland.app</p>
                                    <p><strong>网站：</strong><a href="https://floatisland.app" class="text-indigo-600 hover:underline">https://floatisland.app</a></p>
                                </div>
                            </div>
                        </section>
                        <div class="border-t pt-6 mt-8">
                            <p class="text-sm text-gray-600 text-center">
                                本使用条款适用于浮岛（Float Island）应用及相关服务。<br>
                                最后更新：2025年12月8日
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    `;

    document.querySelector('#app').innerHTML = htmlContent;

    // Initialize app after content is loaded
    initializeApp();
}

// Load the page
loadPage();

// Initialize all functionality after page loads
function initializeApp() {
    // Wait for GSAP to be available
    if (typeof gsap === 'undefined') {
        console.error('GSAP not loaded');
        return;
    }

    gsap.registerPlugin(ScrollTrigger);

    // Download button click handler
     const downloadBtn = document.getElementById('download-btn');
     const downloadUrl = 'http://114.116.248.91:888/uploads/apk/android/1_floatisland_v0.0.1.apk';
     if (downloadBtn) {
         downloadBtn.addEventListener('click', (e) => {
             e.preventDefault();
             window.open(downloadUrl, '_blank');
         });
     }

     // Add click handlers for download links
     const downloadLinks = document.querySelectorAll('.download-link');
     downloadLinks.forEach(link => {
         link.addEventListener('click', (e) => {
             e.preventDefault();
             window.open(downloadUrl, '_blank');
         });
     });

    // Modal functionality
    const privacyLink = document.getElementById('privacy-link');
    const termsLink = document.getElementById('terms-link');
    const privacyModal = document.getElementById('privacy-modal');
    const termsModal = document.getElementById('terms-modal');
    const closePrivacyModal = document.getElementById('close-privacy-modal');
    const closeTermsModal = document.getElementById('close-terms-modal');

    // Show privacy modal
    if (privacyLink && privacyModal) {
        privacyLink.addEventListener('click', (e) => {
            e.preventDefault();
            privacyModal.classList.remove('hidden');
            document.body.style.overflow = 'hidden';
        });
    }

    // Show terms modal
    if (termsLink && termsModal) {
        termsLink.addEventListener('click', (e) => {
            e.preventDefault();
            termsModal.classList.remove('hidden');
            document.body.style.overflow = 'hidden';
        });
    }

    // Close privacy modal
    if (closePrivacyModal && privacyModal) {
        closePrivacyModal.addEventListener('click', () => {
            privacyModal.classList.add('hidden');
            document.body.style.overflow = '';
        });
    }

    // Close terms modal
    if (closeTermsModal && termsModal) {
        closeTermsModal.addEventListener('click', () => {
            termsModal.classList.add('hidden');
            document.body.style.overflow = '';
        });
    }

    // Close modals when clicking outside
    [privacyModal, termsModal].forEach(modal => {
        if (modal) {
            modal.addEventListener('click', (e) => {
                if (e.target === modal) {
                    modal.classList.add('hidden');
                    document.body.style.overflow = '';
                }
            });
        }
    });

    // Close modals with Escape key
    document.addEventListener('keydown', (e) => {
        if (e.key === 'Escape') {
            if (privacyModal) privacyModal.classList.add('hidden');
            if (termsModal) termsModal.classList.add('hidden');
            document.body.style.overflow = '';
        }
    });

    // Hero 3D Tilt Effect
    const tiltContainer = document.getElementById('hero-tilt-container');
    const phone = document.getElementById('hero-phone');

    if (tiltContainer && phone) {
        tiltContainer.addEventListener('mousemove', (e) => {
            const rect = tiltContainer.getBoundingClientRect();
            const x = (e.clientX - rect.left) / rect.width - 0.5;
            const y = (e.clientY - rect.top) / rect.height - 0.5;

            gsap.to(phone, {
                rotationY: x * 15,
                rotationX: -y * 15,
                duration: 0.5,
                ease: 'power2.out',
                transformPerspective: 1000
            });

            // Floating elements parallax
            gsap.to('.float-anim', {
                x: -x * 30,
                y: -y * 30,
                duration: 0.5,
                ease: 'power2.out'
            });
        });

        tiltContainer.addEventListener('mouseleave', () => {
            gsap.to(phone, { rotationY: 0, rotationX: 0, duration: 1, ease: 'elastic.out(1, 0.5)' });
            gsap.to('.float-anim', { x: 0, y: 0, duration: 1 });
        });
    }

    // Scroll-triggered animations for feature blocks
    gsap.utils.toArray('.feature-block').forEach(block => {
        gsap.from(block, {
            scrollTrigger: {
                trigger: block,
                start: "top 80%",
            },
            y: 50,
            opacity: 0,
            duration: 1,
            ease: "power3.out"
        });
    });
}