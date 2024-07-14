<h2>ioutil.ReadAll deprecated</h2>

<a href="https://stackoverflow.com/questions/75206234/for-go-ioutil-readall-ioutil-readfile-ioutil-readdir-deprecated">🔗source</a>

<pre>
ioutil.ReadAll -> io.ReadAll
ioutil.ReadFile -> os.ReadFile
ioutil.ReadDir -> os.ReadDir
// others
ioutil.NopCloser -> io.NopCloser
ioutil.TempDir -> os.MkdirTemp
ioutil.TempFile -> os.CreateTemp
ioutil.WriteFile -> os.WriteFile
</pre>