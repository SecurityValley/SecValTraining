" show line numbers
set number

"It hides buffers instead of closing them. This means that you can have unwritten changes to a file and open a new file using :e, without being forced to write or undo your changes first. Also, undo buffers and marks are preserved while the buffer is open.
set hidden

"utf8 default encoding
set encoding=UTF-8

"create no swapfile
set noswapfile

"scroll offset
set scrolloff=20

"50 entries in histroy
set history=50

"4 spaces as tab
set tabstop=4
set softtabstop=4
set shiftwidth=4
set expandtab

"unix file format
set fileformat=unix

"no backups
set nobackup
set nowritebackup

"2 lines cmd height
set cmdheight=2

"speed up command capture time
set updatetime=500
set shortmess+=c

"line break and smart indent
set lbr
set si

"split below and to the right
set splitright
set splitbelow

"dark background
set background=dark

"enable syntax highlight
syntax on
filetype plugin indent on

" netrw config
let g:netrw_liststyle=3
let g:netrw_banner=0
let g:netrw_browse_split=4
let g:netrw_winsize=20
let g:netrw_altv=1
augroup ProjectDrawer
    autocmd!
    autocmd VimEnter * :Lexplore
augroup END

"Plugins
call plug#begin()
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
Plug 'junegunn/fzf', { 'do': { -> fzf#install() } }
Plug 'junegunn/fzf.vim'
Plug 'fatih/vim-go', { 'do': ':GoUpdateBinaries'  }
Plug 'neoclide/coc.nvim', {'branch': 'release'}
Plug 'morhetz/gruvbox'
Plug 'vim-scripts/delimitMate.vim'
Plug 'iamcco/coc-angular'
call plug#end()


"load gruvbox colorscheme
colorscheme gruvbox

"Airline
let g:airline#extensions#tabline#enabled = 1
let g:airline_theme='wombat'

"FZF
nnoremap<F2> :Rg<CR>
nnoremap<F3> :Files<CR>

"Buffers
nnoremap<F5> :Buffers<CR>
map <C-x> :bn<CR>

"Go 
let g:go_gopls_enabled=1

"coc
" Use K to show documentation in preview window.
nnoremap <silent> K :call ShowDocumentation()<CR>

function! ShowDocumentation()
  if CocAction('hasProvider', 'hover')
    call CocActionAsync('doHover')
  else
    call feedkeys('K', 'in')
  endif
endfunction


" Custom split view behaviour
nnoremap <C-s> :vsp<CR>
nnoremap <C-a> :sp<CR>

map <C-m> :Lexplore<CR>
