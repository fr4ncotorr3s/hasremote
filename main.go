package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	filePaths := []string{
		"C:/Program Files (x86)/AA_v123.exe",
		"C:/Program Files (x86)/AteraAgent.exe",
		"C:/Program Files (x86)/AnyDesk/AnyDesk.exe",
		"C:/Program Files (x86)/GoToAssist/GoToAssist.exe",
		"C:/Program Files (x86)/ReceiverSetup64.exe",
		"C:/Program Files (x86)/JumpDesktopConnect.msi",
		"C:/Program Files (x86)/KaseyaRemoteControlHost.exe",
		"C:/Program Files (x86)/LogMeIn.msi",
		"C:/Program Files (x86)/mRemoteNG/mRemoteNG.exe",
		"C:/Program Files (x86)/BAVideoChat.exe",
		"C:/Program Files (x86)/ngrok.exe",
		"C:/Program Files (x86)/pservice.exe",
		"C:/Program Files (x86)/PCMonitorManager.exe",
		"C:/Program Files (x86)/PCMonitorSrv.exe",
		"C:/Program Files (x86)/pcmontask.exe",
		"C:/Program Files (x86)/Radmin.exe",
		"C:/Program Files (x86)/vncviewer.exe",
		"C:/Program Files (x86)/rfusclient.exe",
		"C:/Program Files (x86)/agent-7.1.7.0.exe",
		"C:/Program Files (x86)/RemotePCDesktop.exe",
		"C:/Program Files (x86)/RemotePC.exe",
		"C:/Program Files (x86)/rsocx.exe",
		"C:/Program Files (x86)/RustDesk.exe",
		"C:/Program Files (x86)/ConnectWiseControl.Client.exe",
		"C:/Program Files (x86)/screenconnect.exe",
		"C:/Program Files (x86)/Supremo.exe",
		"C:/Program Files (x86)/TeamViewer/TeamViewer_1.exe",
		"C:/Program Files (x86)/TeamViewer/TeamViewer.exe",
		"C:/Program Files (x86)/tigervvnc123.exe",
		"C:/Program Files (x86)/vncviewer1.exe",
		"C:/Program Files (x86)/tvnserver.exe",
		"C:/Program Files (x86)/tvnviewer.exe",
		"C:/Program Files (x86)/UltraViewer_Desktop.exe",
		"C:/Program Files (x86)/UltraViewer_Service.exe",
		"C:/Program Files (x86)/vncserver.exe",
		"C:/Program Files (x86)/tv_x64.exe",
		"C:/Program Files (x86)/LogMeIn/LogMeIn.exe",
		"C:/Program Files (x86)/AMMYY/AMMYY.exe",
		"C:/Program Files (x86)/AnyDesk/AnyDesk.exe",
		"C:/Program Files (x86)/ATERA.Networks/ATERA.exe",
		"C:/Program Files (x86)/Splashtop/Splashtop.exe",
		"C:/Program Files (x86)/GoToMyPC/GoToMyPC.exe",
		"C:/Program Files (x86)/IntelliAdmin5/IntelliAdmin.exe",
		"C:/Program Files (x86)/mRemoteNG/mRemoteNG.exe",
		"C:/Program Files (x86)/Radmin/Radmin.exe",
		"C:/Program Files (x86)/ScreenConnect/ScreenConnect.exe",
		"C:/Program Files (x86)/Splashtop/Splashtop.exe",
		"C:/Program Files (x86)/Supremo/Supremo.exe",
		"C:/Program Files (x86)/TeamViewer/TeamViewer.exe",
		"C:/Program Files (x86)/UltraViewer/UltraViewer.exe",
		"C:/Program Files (x86)/RealVNC/RealVNC.exe",
		"C:/Program Files (x86)/TigerVNC/TigerVNC.exe",
		"C:/Program Files (x86)/domotz/domotz.exe",
		"C:/Program Files (x86)/LogMeIn/LogMeIn.exe",
		"C:/Program Files (x86)/RemotePC/RemotePC.exe",
		"C:/Program Files (x86)/ScreenConnect/ScreenConnect.exe",
		"C:/Program Files (x86)/TightVNC/TightVNC.exe",
	}

	filePathMap := make(map[string]struct{})
	for _, filePath := range filePaths {
		filePathMap[filePath] = struct{}{}
	}

	uniqueFilePaths := make([]string, 0, len(filePathMap))
	for filePath := range filePathMap {
		uniqueFilePaths = append(uniqueFilePaths, filePath)
	}

	var wg sync.WaitGroup
	resultChan := make(chan string)
	filesFound := false

	for _, filePath := range uniqueFilePaths {
		wg.Add(1)
		go func(filePath string) {
			defer wg.Done()
			if _, err := os.Stat(filePath); err == nil {
				resultChan <- fmt.Sprintf("File exists: %s", filePath)
				filesFound = true
			}
		}(filePath)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for message := range resultChan {
		fmt.Println(message)
	}

	if !filesFound {
		fmt.Println("No files were found.")
	}
}
