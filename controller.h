#ifndef CONTROLLER_H
#define CONTROLLER_H
#include <memory>
#include <iosfwd>
#include <QStringList>
#include "model.h"

class Controller
{
public:
    static Controller &GetInstance();

    bool LoadFromASCII(std::istream &is);

    bool SaveToASCII(std::ostream &os);

    bool LoadMobility();

    void Recalculate();

    const std::vector<double> &GetTemperature();

    const std::vector<double> &GetElectronsConcentration();

    const std::vector<double> &GetHolesConcentration();

    const std::vector<double> &GetDonorsConcentration();

    const std::vector<double> &GetAcceptorsConcentration();

    const std::vector<double> &GetElectronsMobility();

    const std::vector<double> &GetHolesMobility();

    const std::vector<double> &GetConductivity();

private:
    Controller();

    Model model;

    bool CallMobility(QStringList args, QByteArray &data);
};

#endif // CONTROLLER_H
