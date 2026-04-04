package server
import "net/http"
func (s *Server) dashboard(w http.ResponseWriter, r *http.Request) { w.Header().Set("Content-Type", "text/html"); w.Write([]byte(dashHTML)) }
const dashHTML = `<!DOCTYPE html><html><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1.0"><title>Decision Log</title>
<link href="https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;500;700&display=swap" rel="stylesheet">
<style>:root{--bg:#1a1410;--bg2:#241e18;--bg3:#2e261e;--rust:#e8753a;--leather:#a0845c;--cream:#f0e6d3;--cd:#bfb5a3;--cm:#7a7060;--gold:#d4a843;--green:#4a9e5c;--red:#c94444;--blue:#5b8dd9;--mono:'JetBrains Mono',monospace}*{margin:0;padding:0;box-sizing:border-box}body{background:var(--bg);color:var(--cream);font-family:var(--mono);line-height:1.5}.hdr{padding:1rem 1.5rem;border-bottom:1px solid var(--bg3);display:flex;justify-content:space-between;align-items:center}.hdr h1{font-size:.9rem;letter-spacing:2px}.hdr h1 span{color:var(--rust)}.main{padding:1.5rem;max-width:800px;margin:0 auto}.stats{display:grid;grid-template-columns:repeat(3,1fr);gap:.5rem;margin-bottom:1rem}.st{background:var(--bg2);border:1px solid var(--bg3);padding:.6rem;text-align:center}.st-v{font-size:1.2rem;font-weight:700}.st-l{font-size:.5rem;color:var(--cm);text-transform:uppercase;letter-spacing:1px;margin-top:.15rem}.toolbar{display:flex;gap:.5rem;margin-bottom:1rem;align-items:center}.search{flex:1;padding:.4rem .6rem;background:var(--bg2);border:1px solid var(--bg3);color:var(--cream);font-family:var(--mono);font-size:.7rem}.search:focus{outline:none;border-color:var(--leather)}.dec{background:var(--bg2);border:1px solid var(--bg3);padding:1rem;margin-bottom:.6rem;transition:border-color .2s}.dec:hover{border-color:var(--leather)}.dec-top{display:flex;justify-content:space-between;align-items:flex-start;gap:.5rem}.dec-title{font-size:.88rem;font-weight:700}.dec-context{font-size:.7rem;color:var(--cd);margin-top:.3rem}.dec-outcome{font-size:.72rem;color:var(--green);margin-top:.3rem;padding:.3rem .5rem;border-left:2px solid var(--green);font-weight:500}.dec-meta{font-size:.55rem;color:var(--cm);margin-top:.35rem;display:flex;gap:.5rem;flex-wrap:wrap;align-items:center}.dec-actions{display:flex;gap:.3rem;flex-shrink:0}.badge{font-size:.5rem;padding:.12rem .35rem;text-transform:uppercase;letter-spacing:1px;border:1px solid}.badge.decided{border-color:var(--green);color:var(--green)}.badge.pending{border-color:var(--gold);color:var(--gold)}.badge.revisit{border-color:var(--rust);color:var(--rust)}.btn{font-size:.6rem;padding:.25rem .5rem;cursor:pointer;border:1px solid var(--bg3);background:var(--bg);color:var(--cd);transition:all .2s}.btn:hover{border-color:var(--leather);color:var(--cream)}.btn-p{background:var(--rust);border-color:var(--rust);color:#fff}.btn-sm{font-size:.55rem;padding:.2rem .4rem}.modal-bg{display:none;position:fixed;inset:0;background:rgba(0,0,0,.65);z-index:100;align-items:center;justify-content:center}.modal-bg.open{display:flex}.modal{background:var(--bg2);border:1px solid var(--bg3);padding:1.5rem;width:500px;max-width:92vw;max-height:90vh;overflow-y:auto}.modal h2{font-size:.8rem;margin-bottom:1rem;color:var(--rust);letter-spacing:1px}.fr{margin-bottom:.6rem}.fr label{display:block;font-size:.55rem;color:var(--cm);text-transform:uppercase;letter-spacing:1px;margin-bottom:.2rem}.fr input,.fr select,.fr textarea{width:100%;padding:.4rem .5rem;background:var(--bg);border:1px solid var(--bg3);color:var(--cream);font-family:var(--mono);font-size:.7rem}.fr input:focus,.fr select:focus,.fr textarea:focus{outline:none;border-color:var(--leather)}.row2{display:grid;grid-template-columns:1fr 1fr;gap:.5rem}.acts{display:flex;gap:.4rem;justify-content:flex-end;margin-top:1rem}.empty{text-align:center;padding:3rem;color:var(--cm);font-style:italic;font-size:.75rem}</style></head><body>
<div class="hdr"><h1><span>&#9670;</span> DECISION LOG</h1><button class="btn btn-p" onclick="openForm()">+ Log Decision</button></div>
<div class="main"><div class="stats" id="stats"></div><div class="toolbar"><input class="search" id="search" placeholder="Search decisions..." oninput="render()"></div><div id="list"></div></div>
<div class="modal-bg" id="mbg" onclick="if(event.target===this)closeModal()"><div class="modal" id="mdl"></div></div>
<script>
var A='/api',items=[],editId=null;
async function load(){var r=await fetch(A+'/decisions').then(function(r){return r.json()});items=r.decisions||[];renderStats();render();}
function renderStats(){var total=items.length;var decided=items.filter(function(d){return d.status==='decided'}).length;var pending=items.filter(function(d){return d.status==='pending'}).length;
document.getElementById('stats').innerHTML='<div class="st"><div class="st-v">'+total+'</div><div class="st-l">Decisions</div></div><div class="st"><div class="st-v" style="color:var(--green)">'+decided+'</div><div class="st-l">Decided</div></div><div class="st"><div class="st-v" style="color:var(--gold)">'+pending+'</div><div class="st-l">Pending</div></div>';}
function render(){var q=(document.getElementById('search').value||'').toLowerCase();var f=items;
if(q)f=f.filter(function(d){return(d.title||'').toLowerCase().includes(q)||(d.context||'').toLowerCase().includes(q)||(d.outcome||'').toLowerCase().includes(q)});
if(!f.length){document.getElementById('list').innerHTML='<div class="empty">No decisions logged.</div>';return;}
var h='';f.forEach(function(d){
h+='<div class="dec"><div class="dec-top"><div style="flex:1"><div class="dec-title">'+esc(d.title)+'</div></div>';
h+='<div class="dec-actions"><button class="btn btn-sm" onclick="openEdit(''+d.id+'')">Edit</button><button class="btn btn-sm" onclick="del(''+d.id+'')" style="color:var(--red)">&#10005;</button></div></div>';
if(d.context)h+='<div class="dec-context">'+esc(d.context)+'</div>';
if(d.outcome)h+='<div class="dec-outcome">&#10003; '+esc(d.outcome)+'</div>';
h+='<div class="dec-meta">';
var st=d.status||'pending';h+='<span class="badge '+st+'">'+st+'</span>';
if(d.decided_by)h+='<span>'+esc(d.decided_by)+'</span>';
if(d.decided_at)h+='<span>'+d.decided_at+'</span>';
h+='<span>'+ft(d.created_at)+'</span></div></div>';});
document.getElementById('list').innerHTML=h;}
async function del(id){if(!confirm('Delete?'))return;await fetch(A+'/decisions/'+id,{method:'DELETE'});load();}
function formHTML(dec){var i=dec||{title:'',context:'',options:'',outcome:'',decided_by:'',decided_at:'',status:'pending'};var isEdit=!!dec;
var h='<h2>'+(isEdit?'EDIT':'LOG')+' DECISION</h2>';
h+='<div class="fr"><label>Title *</label><input id="f-title" value="'+esc(i.title)+'" placeholder="What was decided?"></div>';
h+='<div class="fr"><label>Context</label><textarea id="f-context" rows="3" placeholder="Why was this decision needed?">'+esc(i.context)+'</textarea></div>';
h+='<div class="fr"><label>Options Considered</label><textarea id="f-options" rows="2" placeholder="What were the alternatives?">'+esc(i.options)+'</textarea></div>';
h+='<div class="fr"><label>Outcome</label><textarea id="f-outcome" rows="2" placeholder="What was decided and why?">'+esc(i.outcome)+'</textarea></div>';
h+='<div class="row2"><div class="fr"><label>Decided By</label><input id="f-by" value="'+esc(i.decided_by)+'"></div><div class="fr"><label>Status</label><select id="f-status">';
['pending','decided','revisit'].forEach(function(s){h+='<option value="'+s+'"'+(i.status===s?' selected':'')+'>'+s+'</option>';});
h+='</select></div></div>';
h+='<div class="fr"><label>Decided Date</label><input id="f-date" type="date" value="'+esc(i.decided_at)+'"></div>';
h+='<div class="acts"><button class="btn" onclick="closeModal()">Cancel</button><button class="btn btn-p" onclick="submit()">'+(isEdit?'Save':'Log')+'</button></div>';return h;}
function openForm(){editId=null;document.getElementById('mdl').innerHTML=formHTML();document.getElementById('mbg').classList.add('open');}
function openEdit(id){var d=null;for(var j=0;j<items.length;j++){if(items[j].id===id){d=items[j];break;}}if(!d)return;editId=id;document.getElementById('mdl').innerHTML=formHTML(d);document.getElementById('mbg').classList.add('open');}
function closeModal(){document.getElementById('mbg').classList.remove('open');editId=null;}
async function submit(){var title=document.getElementById('f-title').value.trim();if(!title){alert('Title required');return;}
var body={title:title,context:document.getElementById('f-context').value.trim(),options:document.getElementById('f-options').value.trim(),outcome:document.getElementById('f-outcome').value.trim(),decided_by:document.getElementById('f-by').value.trim(),status:document.getElementById('f-status').value,decided_at:document.getElementById('f-date').value};
if(editId){await fetch(A+'/decisions/'+editId,{method:'PUT',headers:{'Content-Type':'application/json'},body:JSON.stringify(body)});}
else{await fetch(A+'/decisions',{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify(body)});}closeModal();load();}
function ft(t){if(!t)return'';try{return new Date(t).toLocaleDateString('en-US',{month:'short',day:'numeric'})}catch(e){return t;}}
function esc(s){if(!s)return'';var d=document.createElement('div');d.textContent=s;return d.innerHTML;}
document.addEventListener('keydown',function(e){if(e.key==='Escape')closeModal();});load();
</script></body></html>`
