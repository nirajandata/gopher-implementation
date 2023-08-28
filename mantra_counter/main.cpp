#include <wx/wx.h>
#include "out.xpm"

class MyFrame : public wxFrame
{
public:
    MyFrame() : wxFrame(NULL, wxID_ANY, "Counter")
    {
        count = 0;

        // Create the value label
        valueLabel = new wxStaticText(this, wxID_ANY, wxString::Format(wxT("%d"), count));
        wxFont font(20, wxFONTFAMILY_DEFAULT, wxFONTSTYLE_NORMAL, wxFONTWEIGHT_NORMAL);
        valueLabel->SetFont(font);

        /* since, we are embedding into binary, we ignore this
        wxString imagePath = wxT("ab.jpg");
        wxBitmap image;
        image.LoadFile("ab.jpg",wxBITMAP_TYPE_JPEG);
        wxBitmap bitmap(image);
         */
        wxBitmap bitmap=wxBitmap(out);

        // Create the image control
        imageControl = new wxStaticBitmap(this, wxID_ANY, bitmap);

        // Create the buttons
        countButton = new wxButton(this, wxID_ANY, wxT("Count"));
        resetButton = new wxButton(this, wxID_ANY, wxT("Reset"));

        // Bind the button events
        countButton->Bind(wxEVT_BUTTON, &MyFrame::OnCountButtonClicked, this);
        resetButton->Bind(wxEVT_BUTTON, &MyFrame::OnResetButtonClicked, this);

        countButton->SetMinSize(wxSize(100, 50));
        resetButton->SetMinSize(wxSize(100, 50));
        // Arrange the elements using a sizer
        wxBoxSizer* sizer = new wxBoxSizer(wxVERTICAL);
        sizer->Add(valueLabel, 0,  wxALL | wxALIGN_CENTER_HORIZONTAL, 10);
        sizer->Add(imageControl, 0, wxALIGN_CENTRE_HORIZONTAL , 10);

        // Create a horizontal sizer for the buttons
        wxBoxSizer* buttonSizer = new wxBoxSizer(wxHORIZONTAL);
        buttonSizer->Add(countButton, 0, wxALL, 10);
        buttonSizer->AddStretchSpacer(); // Add a stretchable space
        buttonSizer->Add(resetButton, 0,  wxALL, 10);

        sizer->Add(buttonSizer, 0, wxEXPAND);
        SetSizerAndFit(sizer);
    }

private:
    void OnCountButtonClicked(wxCommandEvent& event)
    {
        count++;
        valueLabel->SetLabel(wxString::Format(wxT("%d"), count));
    }

    void OnResetButtonClicked(wxCommandEvent& event)
    {
        count = 0;
        valueLabel->SetLabel(wxString::Format(wxT("%d"), count));
    }

    wxStaticText* valueLabel;
    wxStaticBitmap* imageControl;
    wxButton* countButton;
    wxButton* resetButton;
    int count;
};

class MyApp : public wxApp
{
public:
    virtual bool OnInit()
    {
        MyFrame* frame = new MyFrame();
        frame->Show();
        return true;
    }
};

wxIMPLEMENT_APP(MyApp);
