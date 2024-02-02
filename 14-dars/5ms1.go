func firstPalindrome(words []string) string {
	g:=0
	for i:=0;i<len(words);i++{
	   b:=words[i]
   s:=""
   
	   for a:=(len(b)-1);a>=0;a--{
	   s=s+string(b[a])
	   }
	   if s==words[i]&&g==0{
		   v:=("The first string that is palindromic is ",words[i])
		   g+=1
		   return v
	   
	   
	   }
	   if s==words[i]&&g==1{
		   p:=v+("Note that " ,words[i]," is also palindromic, but it is not the first. ")
		   g+=1
		   return p
	   
	   }
	   if s==words[i]&&g==2{
		   o:=p+("Note that " ,words[i]," is also palindromic, but it is not the first.")
		   g+=1
		   return o
	   
	   
	   }
   }   
   