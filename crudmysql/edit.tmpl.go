{{ define "Edit" }}
  {{ template "Header" }}
    {{ template "Menu" }} 
   <h2>Edit Name and City</h2>  
    <form method="POST" action="update">
      <input type="hidden" name="uid" value="{{ .Id }}" />
      <label> Name </label><input type="text" name="name" value="{{ .Name }}"  /><br />
      <label> City </label><input type="text" name="city" value="{{ .City }}"  /><br />
      <input type="submit" value="Save user" />
    </form><br />    
  {{ template "Footer" }}
{{ end }}