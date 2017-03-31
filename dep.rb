require 'rake'

def dep(command, title: nil, error: "the operation failed")
  title = "Command: #{command}" unless title
  if title
    puts "\# #{title}"
  end
  sh command do |ok, res|
     if !ok
       abort error
     end
  end
end

dep "goapp get github.com/gorilla/mux"
dep "goapp get github.com/asaskevich/govalidator"
dep "goapp get gopkg.in/gin-gonic/gin.v1"
dep "goapp get github.com/huandu/facebook"

puts '----- OPERATION COMPLITED ------'
